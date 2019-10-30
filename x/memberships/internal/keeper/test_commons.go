package keeper

import (
	"time"

	"github.com/commercionetwork/commercionetwork/x/government"
	"github.com/commercionetwork/commercionetwork/x/memberships/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/cosmos/modules/incubator/nft"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

//This function create an environment to test modules
func GetTestInput() (*codec.Codec, sdk.Context, bank.Keeper, government.Keeper, Keeper) {

	memDB := db.NewMemDB()
	cdc := testCodec()

	keys := sdk.NewKVStoreKeys(
		auth.StoreKey,
		params.StoreKey,
		supply.StoreKey,
		nft.StoreKey,
		government.StoreKey,

		types.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}
	for _, tkey := range tKeys {
		ms.MountStoreWithDB(tkey, sdk.StoreTypeTransient, memDB)
	}
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	pk := params.NewKeeper(cdc, keys[params.StoreKey], tKeys[params.TStoreKey], params.DefaultCodespace)
	ak := auth.NewAccountKeeper(cdc, keys[auth.StoreKey], pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bk := bank.NewBaseKeeper(ak, pk.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, nil)
	maccPerms := map[string][]string{
		types.ModuleName: {supply.Minter, supply.Burner},
	}
	sk := supply.NewKeeper(cdc, keys[supply.StoreKey], ak, bk, maccPerms)
	sk.SetSupply(ctx, supply.NewSupply(sdk.NewCoins(sdk.NewInt64Coin("stake", 1))))

	govk := government.NewKeeper(cdc, keys[government.StoreKey])
	nftk := nft.NewKeeper(cdc, keys[nft.StoreKey])

	k := NewKeeper(cdc, keys[types.StoreKey], nftk, sk)

	// Set module accounts
	memAcc := supply.NewEmptyModuleAccount(types.ModuleName, supply.Minter, supply.Burner)
	k.supplyKeeper.SetModuleAccount(ctx, memAcc)

	// Set the stable credits denom
	k.SetStableCreditsDenom(ctx, TestStableCreditsDenom)

	return cdc, ctx, bk, govk, k
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	nft.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)

	types.RegisterCodec(cdc)

	cdc.Seal()

	return cdc
}

// Testing variables
var TestUser, _ = sdk.AccAddressFromBech32("cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae")
var TestUser2, _ = sdk.AccAddressFromBech32("cosmos1h7tw92a66gr58pxgmf6cc336lgxadpjz5d5psf")

var TestTsp, _ = sdk.AccAddressFromBech32("cosmos1lwmppctrr6ssnrmuyzu554dzf50apkfvd53jx0")

var zone, _ = time.LoadLocation("UTC")
var TestTimestamp = time.Date(1990, 10, 10, 20, 20, 0, 0, zone)

var TestStableCreditsDenom = "uccc"
var TestMembershipType = "bronze"
