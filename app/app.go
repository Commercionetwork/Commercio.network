package app

import (
	"github.com/cosmos/cosmos-sdk/x/upgrade"

	customUpgrade "github.com/commercionetwork/commercionetwork/x/upgrade"
	customUpgradeKeeper "github.com/commercionetwork/commercionetwork/x/upgrade/keeper"
	customUpgradeTypes "github.com/commercionetwork/commercionetwork/x/upgrade/types"

	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/commercionetwork/commercionetwork/x/ante"
	"github.com/commercionetwork/commercionetwork/x/commerciomint"
	commerciomintKeeper "github.com/commercionetwork/commercionetwork/x/commerciomint/keeper"

	documentsTypes "github.com/commercionetwork/commercionetwork/x/documents/types"

	commerciokycKeeper "github.com/commercionetwork/commercionetwork/x/commerciokyc/keeper"
	documentsKeeper "github.com/commercionetwork/commercionetwork/x/documents/keeper"
	governmentKeeper "github.com/commercionetwork/commercionetwork/x/government/keeper"
	vbrKeeper "github.com/commercionetwork/commercionetwork/x/vbr/keeper"

	commerciokycTypes "github.com/commercionetwork/commercionetwork/x/commerciokyc/types"
	commerciomintTypes "github.com/commercionetwork/commercionetwork/x/commerciomint/types"
	"github.com/commercionetwork/commercionetwork/x/common/types"
	"github.com/commercionetwork/commercionetwork/x/creditrisk"
	creditriskTypes "github.com/commercionetwork/commercionetwork/x/creditrisk/types"
	"github.com/commercionetwork/commercionetwork/x/documents"
	custombank "github.com/commercionetwork/commercionetwork/x/encapsulated/bank"
	customcrisis "github.com/commercionetwork/commercionetwork/x/encapsulated/crisis"
	customstaking "github.com/commercionetwork/commercionetwork/x/encapsulated/staking"
	"github.com/commercionetwork/commercionetwork/x/government"
	governmentTypes "github.com/commercionetwork/commercionetwork/x/government/types"
	"github.com/commercionetwork/commercionetwork/x/id"
	idKeeper "github.com/commercionetwork/commercionetwork/x/id/keeper"
	idTypes "github.com/commercionetwork/commercionetwork/x/id/types"
	pricefeedKeeper "github.com/commercionetwork/commercionetwork/x/pricefeed/keeper"
	pricefeedTypes "github.com/commercionetwork/commercionetwork/x/pricefeed/types"
	vbrTypes "github.com/commercionetwork/commercionetwork/x/vbr/types"

	commerciokyc "github.com/commercionetwork/commercionetwork/x/commerciokyc"
	"github.com/commercionetwork/commercionetwork/x/pricefeed"
	"github.com/commercionetwork/commercionetwork/x/vbr"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"

	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"
)

const (
	appName = "Commercio.network"

	DefaultBondDenom   = "ucommercio"
	StableCreditsDenom = "uccc"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32MainPrefix = "did:com:"

	// PrefixValidator is the prefix for validator keys
	PrefixValidator = "val"
	// PrefixConsensus is the prefix for consensus keys
	PrefixConsensus = "cons"
	// PrefixPublic is the prefix for public keys
	PrefixPublic = "pub"
	// PrefixOperator is the prefix for operator keys
	PrefixOperator = "oper"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32MainPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic
)

// default home directories for expected binaries
var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.cncli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.cnd")

	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
		upgrade.AppModuleBasic{},

		// Encapsulated modules
		customcrisis.NewAppModuleBasic(DefaultBondDenom),
		customstaking.NewAppModuleBasic(DefaultBondDenom),
		custombank.NewAppModuleBasic(bank.AppModuleBasic{}),

		// Custom modules
		documents.AppModuleBasic{},
		government.AppModuleBasic{},
		id.AppModuleBasic{},
		commerciokyc.NewAppModuleBasic(StableCreditsDenom),
		commerciomint.NewAppModuleBasic(),
		pricefeed.AppModuleBasic{},
		vbr.AppModuleBasic{},
		creditrisk.AppModuleBasic{},
		customUpgrade.AppModuleBasic{},
	)

	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		distr.ModuleName:          nil,
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},

		// Custom modules
		commerciomintTypes.ModuleName: {supply.Minter, supply.Burner},
		commerciokycTypes.ModuleName:  {supply.Minter, supply.Burner},
		idTypes.ModuleName:            nil,
		vbrTypes.ModuleName:           {supply.Minter},
		creditriskTypes.ModuleName:    nil,
	}

	allowedModuleReceivers = types.Strings{
		commerciomintTypes.ModuleName,
		commerciokycTypes.ModuleName,
		vbrTypes.ModuleName,
		creditriskTypes.ModuleName,
	}
)

// custom tx codec
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)

	return cdc
}

func SetAddressPrefixes() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.Seal()
}

// Extended ABCI application
type CommercioNetworkApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// sdk keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// sdk keepers
	accountKeeper  auth.AccountKeeper
	bankKeeper     bank.Keeper
	supplyKeeper   supply.Keeper
	sendKeeper     bank.SendKeeper
	stakingKeeper  staking.Keeper
	slashingKeeper slashing.Keeper
	distrKeeper    distr.Keeper
	crisisKeeper   crisis.Keeper
	paramsKeeper   params.Keeper
	upgradeKeeper  upgrade.Keeper

	// Encapsulated modules
	customBankKeeper custombank.Keeper

	// Custom modules
	documentsKeeper     documentsKeeper.Keeper
	governmentKeeper    governmentKeeper.Keeper
	idKeeper            idKeeper.Keeper
	membershipKeeper    commerciokycKeeper.Keeper
	mintKeeper          commerciomintKeeper.Keeper
	priceFeedKeeper     pricefeedKeeper.Keeper
	vbrKeeper           vbrKeeper.Keeper
	creditriskKeeper    creditrisk.Keeper
	customUpgradeKeeper customUpgradeKeeper.Keeper

	mm *module.Manager
}

// NewCommercioNetworkApp returns a reference to an initialized CommercioNetworkApp.
func NewCommercioNetworkApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp)) *CommercioNetworkApp {

	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(
		// Basics
		bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, distr.StoreKey, slashing.StoreKey,
		params.StoreKey,
		upgrade.StoreKey,

		// Encapsulated modules
		custombank.StoreKey,

		// Custom modules
		documentsTypes.StoreKey,
		governmentTypes.StoreKey,
		idTypes.StoreKey,
		commerciokycTypes.StoreKey,
		commerciomintTypes.StoreKey,
		pricefeedTypes.StoreKey,
		vbrTypes.StoreKey,
		creditriskTypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)

	// Here you initialize your application with the store keys it requires
	var app = &CommercioNetworkApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tkeys:          tkeys,
	}

	// init params keeper and subspaces
	app.paramsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tkeys[params.TStoreKey])
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSubspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	crisisSubspace := app.paramsKeeper.Subspace(crisis.DefaultParamspace)

	// add keepers
	app.accountKeeper = auth.NewAccountKeeper(app.cdc, keys[auth.StoreKey], authSubspace, auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper, bankSubspace, app.BlacklistedModuleAccAddrs())
	app.supplyKeeper = supply.NewKeeper(app.cdc, keys[supply.StoreKey], app.accountKeeper, app.bankKeeper, maccPerms)
	stakingKeeper := staking.NewKeeper(
		app.cdc, keys[staking.StoreKey],
		app.supplyKeeper, stakingSubspace,
	)
	app.distrKeeper = distr.NewKeeper(app.cdc, keys[distr.StoreKey], distrSubspace, &stakingKeeper,
		app.supplyKeeper, auth.FeeCollectorName, app.ModuleAccountAddrs())
	app.slashingKeeper = slashing.NewKeeper(
		app.cdc, keys[slashing.StoreKey], &stakingKeeper, slashingSubspace,
	)
	app.crisisKeeper = crisis.NewKeeper(crisisSubspace, invCheckPeriod, app.supplyKeeper, auth.FeeCollectorName)
	app.upgradeKeeper = upgrade.NewKeeper(map[int64]bool{}, keys[upgrade.StoreKey], app.cdc)
	// Update demo here https://github.com/regen-network/gaia/blob/gaia-upgrade-demo/docs/upgrade-demo.md
	for upgradeName, upgradeHandler := range upgrades {
		app.upgradeKeeper.SetUpgradeHandler(upgradeName, upgradeHandler)
	}

	// Encapsulated modules
	app.customBankKeeper = custombank.NewKeeper(app.cdc, app.keys[custombank.StoreKey], app.bankKeeper)

	// Custom modules
	app.governmentKeeper = governmentKeeper.NewKeeper(app.cdc, app.keys[governmentTypes.StoreKey])
	app.membershipKeeper = commerciokycKeeper.NewKeeper(app.cdc, app.keys[commerciokycTypes.StoreKey], app.supplyKeeper, app.bankKeeper, app.governmentKeeper, app.accountKeeper)
	app.documentsKeeper = documentsKeeper.NewKeeper(app.keys[documentsTypes.StoreKey], app.governmentKeeper, app.cdc)
	app.idKeeper = idKeeper.NewKeeper(app.cdc, app.keys[idTypes.StoreKey], app.accountKeeper, app.supplyKeeper)
	app.priceFeedKeeper = pricefeedKeeper.NewKeeper(app.cdc, app.keys[pricefeedTypes.StoreKey], app.governmentKeeper)
	app.vbrKeeper = vbrKeeper.NewKeeper(app.cdc, app.keys[vbrTypes.StoreKey], app.distrKeeper, app.supplyKeeper, app.governmentKeeper)
	app.mintKeeper = commerciomintKeeper.NewKeeper(app.cdc, app.keys[commerciomintTypes.StoreKey], app.supplyKeeper, app.priceFeedKeeper, app.governmentKeeper)
	app.creditriskKeeper = creditrisk.NewKeeper(cdc, app.keys[creditriskTypes.StoreKey], app.supplyKeeper)
	app.customUpgradeKeeper = customUpgradeKeeper.NewKeeper(cdc, app.governmentKeeper, app.upgradeKeeper)

	// register the proposal types
	// govRouter := gov.NewRouter()
	// govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler)
	// app.govKeeper = gov.NewKeeper(
	// 	app.cdc, keys[gov.StoreKey], govSubspace,
	// 	app.supplyKeeper, &stakingKeeper, govRouter,
	// )

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.distrKeeper.Hooks(), app.slashingKeeper.Hooks()),
	)

	// Create default modules to be used from customs during encapsulation
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distr.NewAppModule(app.distrKeeper, app.accountKeeper, app.supplyKeeper, app.stakingKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),
		crisis.NewAppModule(&app.crisisKeeper),

		// Encapsulating modules
		custombank.NewAppModule(bank.NewAppModule(app.bankKeeper, app.accountKeeper), app.customBankKeeper, app.governmentKeeper),

		// Custom modules
		documents.NewAppModule(app.documentsKeeper),
		government.NewAppModule(app.governmentKeeper),
		id.NewAppModule(app.idKeeper, app.governmentKeeper, app.supplyKeeper),
		commerciokyc.NewAppModule(app.membershipKeeper, app.supplyKeeper, app.governmentKeeper, app.accountKeeper),
		commerciomint.NewAppModule(app.mintKeeper, app.supplyKeeper),
		pricefeed.NewAppModule(app.priceFeedKeeper, app.governmentKeeper),
		vbr.NewAppModule(app.vbrKeeper, app.stakingKeeper),
		creditrisk.NewAppModule(app.creditriskKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		customUpgrade.NewAppModule(app.customUpgradeKeeper, app.bankKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(
		customUpgradeTypes.ModuleName,
		distr.ModuleName, slashing.ModuleName,
		commerciokycTypes.ModuleName,

		// Custom modules
		vbrTypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisis.ModuleName,
		staking.ModuleName,

		// Custom modules
		pricefeedTypes.ModuleName,
		commerciokycTypes.ModuleName,
		commerciomintTypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		distr.ModuleName, staking.ModuleName, auth.ModuleName, bank.ModuleName,
		slashing.ModuleName, supply.ModuleName,
		commerciokycTypes.ModuleName, crisis.ModuleName, genutil.ModuleName,

		// Custom modules
		governmentTypes.ModuleName,
		documentsTypes.ModuleName,
		idTypes.ModuleName,
		commerciomintTypes.ModuleName,
		pricefeedTypes.ModuleName,
		vbrTypes.ModuleName,
		creditriskTypes.ModuleName,
		upgrade.ModuleName,
	)

	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			app.accountKeeper, app.supplyKeeper, app.priceFeedKeeper, app.governmentKeeper,
			auth.DefaultSigVerificationGasConsumer, StableCreditsDenom,
		),
	)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// application updates every begin block
func (app *CommercioNetworkApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// application updates every end block
func (app *CommercioNetworkApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// application update at chain initialization
func (app *CommercioNetworkApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState simapp.GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.mm.InitGenesis(ctx, genesisState)
}

// load a particular height
func (app *CommercioNetworkApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *CommercioNetworkApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[app.supplyKeeper.GetModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlacklistedModuleAccAddrs returns all the app's module account addresses that
// are black listed from received tokens from the users.
func (app *CommercioNetworkApp) BlacklistedModuleAccAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[app.supplyKeeper.GetModuleAddress(acc).String()] = allowedModuleReceivers.Contains(acc)
	}

	return modAccAddrs
}
