package id

import (
	"github.com/commercionetwork/commercionetwork/x/id/internal/keeper"
	"github.com/commercionetwork/commercionetwork/x/id/internal/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	//function aliases
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	TestSetup     = keeper.SetupTestInput
	RegisterCodec = types.RegisterCodec

	//variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper         = keeper.Keeper
	MsgSetIdentity = types.MsgSetIdentity
)
