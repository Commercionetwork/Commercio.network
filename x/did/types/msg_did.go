package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	commons "github.com/commercionetwork/commercionetwork/x/common/types"
)

func NewMsgSetDidDocument(context string, ID string) *MsgSetDidDocument {
	return &MsgSetDidDocument{
		&DidDocument{Context: []string{context}, ID: ID},
	}

}

func (msg *MsgSetDidDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.DidDocument.ID)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetDidDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetDidDocument) Route() string {
	return RouterKey
}

func (msg *MsgSetDidDocument) Type() string {
	return MsgTypeSetDid
}

func (msg *MsgSetDidDocument) ValidateBasic() error {

	ddo := msg.DidDocument
	// validate ID
	_, err := sdk.AccAddressFromBech32(ddo.ID)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	// validate Context
	if commons.Strings(ddo.Context).Contains(ContextDidV1) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid context, must include %s", ContextDidV1)
	}

	// validate VerificationMethod
	// for _, vm := range ddo.VerificationMethod{
	// 	if vm.ID == "" || vm == VerificationMethod{}
	// }

	// validate Service
	for _, s := range ddo.Service {
		err = s.Validate()
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid service %s %e", s, err)
		}
	}

	// validate Authentication

	// validate AssertionMethod

	// validate CapabilityDelegation

	// validate CapabilityInvocation

	// validate KeyAgreement

	// controller, _ := sdk.AccAddressFromBech32(msg.ID)

	// for _, key := range msg.PubKeys {
	// 	if err := key.Validate(); err != nil {
	// 		return err
	// 	}
	// 	keycontroller, _ := sdk.AccAddressFromBech32(key.Controller)
	// 	if !controller.Equals(keycontroller) {
	// 		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Public key controller must match did document id")
	// 	}
	// }

	// var pubKeys PubKeys
	// pubKeys = msg.PubKeys
	// if err := pubKeys.noDuplicates(); err != nil {
	// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	// }

	// if !pubKeys.HasVerificationAndSignatureKey() {
	// 	return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "specified public keys are not in the correct format")
	// }
	// /*
	// 	if msg.Proof == nil {
	// 		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "proof not provided")
	// 	}

	// 	if err := msg.Proof.Validate(); err != nil {
	// 		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("proof validation error: %s", err.Error()))
	// 	}
	// */
	// // we have some service, we should validate 'em
	// if msg.Service != nil {
	// 	for i, service := range msg.Service {
	// 		err := service.Validate()
	// 		if err != nil {
	// 			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("service %d validation failed: %s", i, err.Error()))
	// 		}
	// 	}
	// }
	// /*
	// 	if err := msg.VerifyProof(); err != nil {
	// 		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	// 	}*/

	// if err := msg.lengthLimits(); err != nil {
	// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	// }

	return nil
}
