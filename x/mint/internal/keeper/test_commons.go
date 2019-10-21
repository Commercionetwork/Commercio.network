package keeper

import (
	"time"

	"github.com/commercionetwork/commercionetwork/x/government"
	"github.com/commercionetwork/commercionetwork/x/mint/internal/types"
	"github.com/commercionetwork/commercionetwork/x/pricefeed"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

var TestCreditsDenom = "stake"
var TestLiquidityDenom = "ucommercio"
var TestOwner, _ = sdk.AccAddressFromBech32("cosmos1lwmppctrr6ssnrmuyzu554dzf50apkfvd53jx0")
var timezone, _ = time.LoadLocation("UTC")
var TestTimestamp = time.Date(1990, 01, 01, 20, 20, 00, 0, timezone)

var TestCdpRequest = types.CdpRequest{
	Signer:          TestOwner,
	DepositedAmount: sdk.NewCoins(sdk.NewCoin(TestLiquidityDenom, sdk.NewInt(100))),
	Timestamp:       TestTimestamp,
}

var TestCdp = types.Cdp{
	Owner:           TestOwner,
	DepositedAmount: sdk.NewCoins(sdk.NewCoin(TestLiquidityDenom, sdk.NewInt(100))),
	CreditsAmount:   sdk.NewCoins(sdk.NewCoin(TestCreditsDenom, sdk.NewInt(50))),
	Timestamp:       TestTimestamp,
}

var TestLiquidityPool = sdk.Coins{sdk.NewInt64Coin(TestLiquidityDenom, 10000)}

func SetupTestInput() (*codec.Codec, sdk.Context, bank.Keeper, pricefeed.Keeper, Keeper) {
	memDB := db.NewMemDB()
	cdc := testCodec()

	authKey := sdk.NewKVStoreKey("authCapKey")
	ibcKey := sdk.NewKVStoreKey("ibcCapKey")
	fckCapKey := sdk.NewKVStoreKey("fckCapKey")
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	tkeyStaking := sdk.NewTransientStoreKey(staking.TStoreKey)
	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
	distrKey := sdk.NewKVStoreKey("distrKey")

	//custom modules keys
	govKey := sdk.NewKVStoreKey(government.StoreKey)
	pricefeedKey := sdk.NewKVStoreKey(pricefeed.StoreKey)
	cMintKey := sdk.NewKVStoreKey(types.StoreKey)

	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(ibcKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(fckCapKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, memDB)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(tkeyStaking, sdk.StoreTypeTransient, nil)
	ms.MountStoreWithDB(keyStaking, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(distrKey, sdk.StoreTypeIAVL, memDB)

	//custom modules loading
	ms.MountStoreWithDB(pricefeedKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(cMintKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(govKey, sdk.StoreTypeIAVL, memDB)

	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	pk := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	ak := auth.NewAccountKeeper(cdc, authKey, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bk := bank.NewBaseKeeper(ak, pk.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, nil)

	pricefeedK := pricefeed.NewKeeper(cdc, pricefeedKey)

	mintK := NewKeeper(cMintKey, bk, pricefeedK, cdc)
	mintK.SetCreditsDenom(ctx, TestCreditsDenom)

	return cdc, ctx, bk, pricefeedK, mintK
}

func testCodec() *codec.Codec {
	var cdc = codec.New()
	sdk.RegisterCodec(cdc)

	bank.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)

	//custom modules codecs
	pricefeed.RegisterCodec(cdc)
	government.RegisterCodec(cdc)

	codec.RegisterCrypto(cdc)

	cdc.Seal()

	return cdc
}
