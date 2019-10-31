package keeper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/commercionetwork/commercionetwork/x/id/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stretchr/testify/assert"
)

func TestValidMsg_StoreDoc(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	handler := NewHandler(k, govK)
	msgSetIdentity := types.MsgSetIdentity(TestDidDocument)
	res := handler(ctx, msgSetIdentity)

	assert.True(t, res.IsOK())
}

func TestInvalidMsg(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	handler := NewHandler(k, govK)
	res := handler(ctx, sdk.NewTestMsg())

	assert.False(t, res.IsOK())
	assert.True(t, strings.Contains(res.Log, fmt.Sprintf("Unrecognized %s message type", types.ModuleName)))
}

// ----------------------------
// --- Did deposit requests
// ----------------------------

var msgRequestDidDeposit = types.MsgRequestDidDeposit{
	Recipient:     TestDidDepositRequest.Recipient,
	Amount:        TestDidDepositRequest.Amount,
	Proof:         TestDidDepositRequest.Proof,
	EncryptionKey: TestDidDepositRequest.EncryptionKey,
	FromAddress:   TestDidDepositRequest.FromAddress,
}

func Test_handleMsgRequestDidDeposit_NewRequest(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	handler := NewHandler(k, govK)
	res := handler(ctx, msgRequestDidDeposit)
	assert.True(t, res.IsOK())

	stored, found := k.GetDidDepositRequestByProof(ctx, TestDidDepositRequest.Proof)
	assert.True(t, found)
	assert.Equal(t, TestDidDepositRequest, stored)
}

func Test_handleMsgRequestDidDeposit_ExistingRequest(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	handler := NewHandler(k, govK)
	res := handler(ctx, msgRequestDidDeposit)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
}

func Test_handleMsgChangeDidDepositRequestStatus_Approved_ReturnsError(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	status := types.RequestStatus{Type: types.StatusApproved, Message: ""}
	msg := types.NewMsgInvalidateDidDepositRequest(status, TestDidDepositRequest.Proof, TestDidDepositRequest.Recipient)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
	assert.Contains(t, res.Log, msg.Status.Type)
}

func Test_handleMsgChangeDidDepositRequestStatus_Rejected_InvalidGovernment(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	status := types.RequestStatus{Type: types.StatusRejected, Message: ""}
	msg := types.NewMsgInvalidateDidDepositRequest(status, TestDidDepositRequest.Proof, TestDidDepositRequest.Recipient)
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeInvalidAddress, res.Code)
	assert.Contains(t, res.Log, msg.Status.Type)
}

func Test_handleMsgChangeDidDepositRequestStatus_Rejected_ValidGovernment(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	status := types.RequestStatus{Type: types.StatusRejected, Message: ""}
	msg := types.NewMsgInvalidateDidDepositRequest(status, TestDidDepositRequest.Proof, govK.GetGovernmentAddress(ctx))

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.True(t, res.IsOK())
}

func Test_handleMsgChangeDidDepositRequestStatus_Canceled_InvalidAddress(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	addr, _ := sdk.AccAddressFromBech32("cosmos1elzra8xnfqhqg2dh5ae9x45tnmud5wazkp92r9")
	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidDepositRequest(status, TestDidDepositRequest.Proof, addr)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeInvalidAddress, res.Code)
	assert.Contains(t, res.Log, "poster")
}

func Test_handleMsgChangeDidDepositRequestStatus_Cancel_ValidAddress(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidDepositRequest(status, TestDidDepositRequest.Proof, TestDidDepositRequest.FromAddress)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.True(t, res.IsOK())

	stored, found := k.GetDidDepositRequestByProof(ctx, TestDidDepositRequest.Proof)
	assert.True(t, found)
	assert.Equal(t, status, *stored.Status)
}

func Test_handleMsgChangeDidDepositRequestStatus_StatusAlreadySet(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	request := types.DidDepositRequest{
		Status:        &types.RequestStatus{Type: types.StatusApproved, Message: ""},
		Recipient:     TestDidDepositRequest.Recipient,
		Amount:        TestDidDepositRequest.Amount,
		Proof:         TestDidDepositRequest.Proof,
		EncryptionKey: TestDidDepositRequest.EncryptionKey,
		FromAddress:   TestDidDepositRequest.FromAddress,
	}
	_ = k.StoreDidDepositRequest(ctx, request)

	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidDepositRequest(status, TestDidDepositRequest.Proof, TestDidDepositRequest.FromAddress)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
	assert.Contains(t, res.Log, "status")
}

func Test_handleMsgChangeDidDepositRequestStatus_AllGood(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidDepositRequest(status, TestDidDepositRequest.Proof, TestDidDepositRequest.FromAddress)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.True(t, res.IsOK())

	stored, found := k.GetDidDepositRequestByProof(ctx, TestDidDepositRequest.Proof)
	assert.True(t, found)
	assert.Equal(t, status, *stored.Status)
}

// ----------------------------
// --- Did power up requests
// --------------------------

var msgRequestDidPowerUp = types.MsgRequestDidPowerUp{
	Claimant:      TestDidPowerUpRequest.Claimant,
	Amount:        TestDidPowerUpRequest.Amount,
	Proof:         TestDidPowerUpRequest.Proof,
	EncryptionKey: TestDidPowerUpRequest.EncryptionKey,
}

func Test_handleMsgRequestDidPowerUp_NewRequest(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	handler := NewHandler(k, govK)
	res := handler(ctx, msgRequestDidPowerUp)
	assert.True(t, res.IsOK())

	stored, found := k.GetPowerUpRequestByProof(ctx, TestDidPowerUpRequest.Proof)
	assert.True(t, found)
	assert.Equal(t, TestDidPowerUpRequest, stored)
}

func Test_handleMsgRequestDidPowerUp_ExistingRequest(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StorePowerUpRequest(ctx, TestDidPowerUpRequest)

	handler := NewHandler(k, govK)
	res := handler(ctx, msgRequestDidPowerUp)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
}

func Test_handleMsgChangeDidPowerUpRequestStatus_Approved_ReturnsError(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StorePowerUpRequest(ctx, TestDidPowerUpRequest)

	status := types.RequestStatus{Type: types.StatusApproved, Message: ""}
	msg := types.NewMsgInvalidateDidPowerUpRequest(status, TestDidPowerUpRequest.Proof, TestDidPowerUpRequest.Claimant)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
	assert.Contains(t, res.Log, msg.Status.Type)
}

func Test_handleMsgChangeDidPowerUpRequestStatus_Rejected_InvalidGovernment(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StorePowerUpRequest(ctx, TestDidPowerUpRequest)

	status := types.RequestStatus{Type: types.StatusRejected, Message: ""}
	msg := types.NewMsgInvalidateDidPowerUpRequest(status, TestDidPowerUpRequest.Proof, TestDidPowerUpRequest.Claimant)
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeInvalidAddress, res.Code)
	assert.Contains(t, res.Log, msg.Status.Type)
}

func Test_handleMsgChangeDidPowerUpRequestStatus_Rejected_ValidGovernment(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StorePowerUpRequest(ctx, TestDidPowerUpRequest)

	status := types.RequestStatus{Type: types.StatusRejected, Message: ""}
	msg := types.NewMsgInvalidateDidPowerUpRequest(status, TestDidPowerUpRequest.Proof, govK.GetGovernmentAddress(ctx))

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.True(t, res.IsOK())
}

func Test_handleMsgChangeDidPowerUpRequestStatus_Canceled_InvalidAddress(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StorePowerUpRequest(ctx, TestDidPowerUpRequest)

	addr, _ := sdk.AccAddressFromBech32("cosmos1elzra8xnfqhqg2dh5ae9x45tnmud5wazkp92r9")
	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidPowerUpRequest(status, TestDidPowerUpRequest.Proof, addr)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeInvalidAddress, res.Code)
	assert.Contains(t, res.Log, "poster")
}

func Test_handleMsgChangeDidPowerUpRequestStatus_Cancel_ValidAddress(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StorePowerUpRequest(ctx, TestDidPowerUpRequest)

	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidPowerUpRequest(status, TestDidPowerUpRequest.Proof, TestDidPowerUpRequest.Claimant)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.True(t, res.IsOK())

	stored, found := k.GetPowerUpRequestByProof(ctx, TestDidPowerUpRequest.Proof)
	assert.True(t, found)
	assert.Equal(t, status, *stored.Status)
}

func Test_handleMsgChangeDidPowerUpRequestStatus_StatusAlreadySet(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	request := types.DidPowerUpRequest{
		Status:        &types.RequestStatus{Type: types.StatusApproved, Message: ""},
		Amount:        TestDidPowerUpRequest.Amount,
		Proof:         TestDidPowerUpRequest.Proof,
		EncryptionKey: TestDidPowerUpRequest.EncryptionKey,
		Claimant:      TestDidPowerUpRequest.Claimant,
	}
	_ = k.StorePowerUpRequest(ctx, request)

	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidPowerUpRequest(status, TestDidPowerUpRequest.Proof, TestDidPowerUpRequest.Claimant)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
	assert.Contains(t, res.Log, "status")
}

func Test_handleMsgChangeDidPowerUpRequestStatus_AllGood(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StorePowerUpRequest(ctx, TestDidPowerUpRequest)

	status := types.RequestStatus{Type: types.StatusCanceled, Message: ""}
	msg := types.NewMsgInvalidateDidPowerUpRequest(status, TestDidPowerUpRequest.Proof, TestDidPowerUpRequest.Claimant)

	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.True(t, res.IsOK())

	stored, found := k.GetPowerUpRequestByProof(ctx, TestDidPowerUpRequest.Proof)
	assert.True(t, found)
	assert.Equal(t, status, *stored.Status)
}

// ------------------------
// --- Deposits handling
// ------------------------

func Test_handleMsgWithdrawDeposit_InvalidGovernment(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)

	msg := types.NewMsgMoveDeposit("", TestDidDepositRequest.FromAddress)
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeInvalidAddress, res.Code)
	assert.Contains(t, res.Log, "government")
}

func Test_handleMsgWithdrawDeposit_InvalidRequestProof(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	msg := types.NewMsgMoveDeposit("", govK.GetGovernmentAddress(ctx))
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
	assert.Contains(t, res.Log, "not found")
}

func Test_handleMsgWithdrawDeposit_RequestAlreadyHasAStatus(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	request := types.DidDepositRequest{
		Status: &types.RequestStatus{
			Type:    "accepted",
			Message: "",
		},
		Recipient:     TestDidDepositRequest.Recipient,
		Amount:        TestDidDepositRequest.Amount,
		Proof:         TestDidDepositRequest.Proof,
		EncryptionKey: TestDidDepositRequest.EncryptionKey,
		FromAddress:   TestDidDepositRequest.FromAddress,
	}
	_ = k.StoreDidDepositRequest(ctx, request)

	msg := types.NewMsgMoveDeposit(request.Proof, govK.GetGovernmentAddress(ctx))
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
	assert.Contains(t, res.Log, "already has a valid status")
}

func Test_handleMsgWithdrawDeposit_AllGood(t *testing.T) {
	_, ctx, _, bK, govK, k := SetupTestInput()
	_ = k.StoreDidDepositRequest(ctx, TestDidDepositRequest)
	_ = bK.SetCoins(ctx, TestDidDepositRequest.FromAddress, TestDidDepositRequest.Amount)

	msg := types.NewMsgMoveDeposit(TestDidDepositRequest.Proof, govK.GetGovernmentAddress(ctx))
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)
	assert.True(t, res.IsOK())

	// Check the balances
	assert.Equal(t, TestDidDepositRequest.Amount, k.GetPoolAmount(ctx))
	assert.Empty(t, bK.GetCoins(ctx, TestDidDepositRequest.FromAddress))
	assert.Empty(t, bK.GetCoins(ctx, TestDidDepositRequest.Recipient))

	// Check the request
	request, _ := k.GetDidDepositRequestByProof(ctx, TestDidDepositRequest.Proof)
	assert.NotNil(t, request.Status)
	assert.Equal(t, types.StatusApproved, request.Status.Type)
}

func Test_handleMsgPowerUpDid_InvalidGovernment(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	msg := types.MsgPowerUpDid{
		Recipient:           TestDidPowerUpRequest.Claimant,
		Amount:              TestDidPowerUpRequest.Amount,
		ActivationReference: "xxxxxx",
		Signer:              TestDidPowerUpRequest.Claimant,
	}
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeInvalidAddress, res.Code)
	assert.Contains(t, res.Log, "government")
}

func Test_handleMsgPowerUpDid_ReferenceAlreadyPresent(t *testing.T) {
	_, ctx, _, _, govK, k := SetupTestInput()

	reference := "xxxxxx"
	k.SetHandledPowerUpRequestsReferences(ctx, []string{reference})

	msg := types.MsgPowerUpDid{
		Recipient:           TestDidPowerUpRequest.Claimant,
		Amount:              TestDidPowerUpRequest.Amount,
		ActivationReference: reference,
		Signer:              govK.GetGovernmentAddress(ctx),
	}
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.False(t, res.IsOK())
	assert.Equal(t, sdk.CodeUnknownRequest, res.Code)
	assert.Contains(t, res.Log, "already handled")
}

func Test_handleMsgPowerUpDid_AllGood(t *testing.T) {
	_, ctx, _, bK, govK, k := SetupTestInput()

	msg := types.MsgPowerUpDid{
		Recipient:           TestDidPowerUpRequest.Claimant,
		Amount:              TestDidPowerUpRequest.Amount,
		ActivationReference: "test-reference",
		Signer:              govK.GetGovernmentAddress(ctx),
	}

	k.supplyKeeper.SetSupply(ctx, supply.NewSupply(msg.Amount))
	_ = bK.SetCoins(ctx, k.supplyKeeper.GetModuleAddress(types.ModuleName), msg.Amount)
	handler := NewHandler(k, govK)
	res := handler(ctx, msg)

	assert.True(t, res.IsOK())

	// Check the balances
	assert.Equal(t, msg.Amount, bK.GetCoins(ctx, msg.Recipient))
	assert.Empty(t, k.GetPoolAmount(ctx))

	// Check the request
	assert.True(t, k.GetHandledPowerUpRequestsReferences(ctx).Contains(msg.ActivationReference))
}
