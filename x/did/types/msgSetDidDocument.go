package types

import (
	commons "github.com/commercionetwork/commercionetwork/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (msg *MsgSetDidDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ID)
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

	// validate ID
	if err := isValidDidCom(msg.ID); err != nil {
		return err
	}

	// validate @context
	// @context The JSON-LD Context is either a string or a list containing any combination of strings and/or ordered maps.
	// The serialized value of @context MUST be the JSON String https://www.w3.org/ns/did/v1, or a JSON Array where the first item is the JSON String https://www.w3.org/ns/did/v1 and the subsequent items are serialized according to the JSON representation production rules.
	var context commons.Strings = msg.Context
	if !context.Contains(ContextDidV1) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid @context, must include %s", ContextDidV1)
	}

	// validate VerificationMethod
	// TODO check no duplicates id
	for _, vm := range msg.VerificationMethod {
		if err := vm.Validate(); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid VerificationMethod %s %e", vm, err)
		}
	}

	// validate service
	// OPTIONAL
	// If present, the associated value MUST be a set of services, where each service is described by a map.
	// A conforming producer MUST NOT produce multiple service entries with the same id.
	for _, s := range msg.Service {
		if err := s.Validate(); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid Service %s %e", s, err)
		}
	}
	if ServiceSlice(msg.Service).hasDuplicate() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid Service %s found services with the same ID", msg.Service)
	}

	// validate authentication
	for _, a := range msg.Authentication {
		if !msg.HasVerificationMethod(a) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid authentication: %s is not among the verification methods", a)
		}
	}

	// validate assertionMethod
	for _, am := range msg.AssertionMethod {
		if !msg.HasVerificationMethod(am) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid assertionMethod: %s is not among the verification methods", am)
		}
	}

	// validate capabilityDelegation
	for _, cd := range msg.CapabilityDelegation {
		if !msg.HasVerificationMethod(cd) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid capabilityDelegation: %s is not among the verification methods", cd)
		}
	}

	// validate capabilityInvocation
	for _, ci := range msg.CapabilityInvocation {
		if !msg.HasVerificationMethod(ci) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid capabilityInvocation: %s is not among the verification methods", ci)
		}
	}

	// validate keyAgreement
	for _, ka := range msg.KeyAgreement {
		if !msg.HasVerificationMethod(ka) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid keyAgreement: %s is not among the verification methods", ka)
		}
	}

	return nil
}

func (msg MsgSetDidDocument) HasVerificationMethod(id string) bool {
	for _, vm := range msg.VerificationMethod {
		if id == vm.ID {
			return true
		}
	}
	return false
}
