package memberships

import (
	"encoding/json"

	governmentKeeper "github.com/commercionetwork/commercionetwork/x/government/keeper"

	"github.com/commercionetwork/commercionetwork/x/memberships/keeper"
	"github.com/commercionetwork/commercionetwork/x/memberships/types"

	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/commercionetwork/commercionetwork/x/memberships/client/cli"
	"github.com/commercionetwork/commercionetwork/x/memberships/client/rest"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the docs module.
type AppModuleBasic struct {
	stableCreditsDenom string
}

func NewAppModuleBasic(stableCreditsDenom string) AppModuleBasic {
	return AppModuleBasic{
		stableCreditsDenom: stableCreditsDenom,
	}
}

var _ module.AppModuleBasic = AppModuleBasic{}

// module name
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// register module codec
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

// default genesis state
func (amb AppModuleBasic) DefaultGenesis() json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(DefaultGenesisState(amb.stableCreditsDenom))
}

// module genesis validation
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := types.ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

// register rest routes
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

// get the root tx command of this module
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(cdc)
}

// get the root query command of this module
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(cdc)
}

//____________________________________________________________________________

// AppModuleSimulation defines the module simulation functions used by the auth module.
type AppModuleSimulation struct{}

// RegisterStoreDecoder registers a decoder for auth module's types
func (AppModuleSimulation) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {}

//____________________________________________________________________________

// AppModule implements an application module for the id module.
type AppModule struct {
	AppModuleBasic
	AppModuleSimulation
	keeper           keeper.Keeper
	governmentKeeper governmentKeeper.Keeper
	supplyKeeper     supply.Keeper
	accountKeeper    auth.AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper keeper.Keeper, supplyKeeper supply.Keeper, govK governmentKeeper.Keeper, accountKeeper auth.AccountKeeper) AppModule {
	return AppModule{
		AppModuleBasic:      AppModuleBasic{},
		AppModuleSimulation: AppModuleSimulation{},
		keeper:              keeper,
		governmentKeeper:    govK,
		supplyKeeper:        supplyKeeper,
		accountKeeper:       accountKeeper,
	}
}

// module name
func (AppModule) Name() string {
	return types.ModuleName
}

// register invariants
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, am.keeper)
}

// module message route name
func (AppModule) Route() string {
	return types.ModuleName
}

// module handler
func (am AppModule) NewHandler() sdk.Handler {
	return keeper.NewHandler(am.keeper, am.governmentKeeper)
}

// module querier route name
func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// module querier
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return keeper.NewQuerier(am.keeper)
}

// module init-genesis
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, am.supplyKeeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// module export genesis
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return types.ModuleCdc.MustMarshalJSON(gs)
}

// module begin-block
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
}

// module end-block
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
