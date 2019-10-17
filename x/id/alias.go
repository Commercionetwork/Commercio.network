package id

import (
	"github.com/commercionetwork/commercionetwork/x/id/internal/keeper"
	"github.com/commercionetwork/commercionetwork/x/id/internal/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute

	StatusApproved = types.StatusApproved
	StatusRejected = types.StatusRejected
	StatusCanceled = types.StatusCanceled
)

var (
	NewKeeper     = keeper.NewKeeper
	NewHandler    = keeper.NewHandler
	NewQuerier    = keeper.NewQuerier
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc

	NewMsgSetIdentity = types.NewMsgSetIdentity

	NewMsgRequestDidDeposit           = types.NewMsgRequestDidDeposit
	NewMsgInvalidateDidDepositRequest = types.NewMsgInvalidateDidDepositRequest
	NewMsgRequestDidPowerUp           = types.NewMsgRequestDidPowerUp
	NewMsgInvalidateDidPowerUpRequest = types.NewMsgInvalidateDidPowerUpRequest
	NewMsgMoveDeposit                 = types.NewMsgMoveDeposit
)

type (
	Keeper = keeper.Keeper

	DidDocument       = types.DidDocument
	MsgSetIdentity    = types.MsgSetIdentity
	DidDepositRequest = types.DidDepositRequest
	DidPowerUpRequest = types.DidPowerUpRequest
	RequestStatus     = types.RequestStatus

	// ---------------
	// --- Messages
	// ---------------

	MsgSetIdentity                 = types.MsgSetIdentity
	MsgRequestDidDeposit           = types.MsgRequestDidDeposit
	MsgInvalidateDidDepositRequest = types.MsgInvalidateDidDepositRequest
	MsgRequestDidPowerUp           = types.MsgRequestDidPowerUp
	MsgInvalidateDidPowerUpRequest = types.MsgInvalidateDidPowerUpRequest
	MsgMoveDeposit                 = types.MsgMoveDeposit
	MsgPowerUpDid                  = types.MsgPowerUpDid
)
