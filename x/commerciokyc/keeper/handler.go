package keeper

import (
	"fmt"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/commercionetwork/commercionetwork/x/commerciokyc/types"
	ctypes "github.com/commercionetwork/commercionetwork/x/common/types"
	government "github.com/commercionetwork/commercionetwork/x/government/keeper"
)

// NewHandler is essentially a sub-router that directs messages coming into this module to the proper handler.
func NewHandler(keeper Keeper, governmentKeeper government.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgInviteUser:
			return handleMsgInviteUser(ctx, keeper, msg)
		case types.MsgDepositIntoLiquidityPool:
			return handleMsgDepositIntoPool(ctx, keeper, msg)
		case types.MsgAddTsp:
			return handleMsgAddTrustedSigner(ctx, keeper, governmentKeeper, msg)
		case types.MsgRemoveTsp:
			return handleMsgRemoveTrustedSigner(ctx, keeper, governmentKeeper, msg)
		case types.MsgBuyMembership:
			return handleMsgBuyMembership(ctx, keeper, msg)
		case types.MsgRemoveMembership:
			return handleMsgRemoveMembership(ctx, keeper, msg)
		case types.MsgSetMembership:
			return handleMsgSetMembership(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized %s message type: %v", types.ModuleName, msg.Type())
			return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgInviteUser(ctx sdk.Context, keeper Keeper, msg types.MsgInviteUser) (*sdk.Result, error) {
	if keeper.accountKeeper.GetAccount(ctx, msg.Recipient) != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnauthorized, "cannot invite existing user")
	}

	// Verify that the user that is inviting has already a membership
	if _, err := keeper.GetMembership(ctx, msg.Sender); err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnauthorized, "Cannot send an invitation without having a membership")
	}

	// Try inviting the user
	if err := keeper.InviteUser(ctx, msg.Recipient, msg.Sender); err != nil {
		return nil, err
	}
	ctypes.EmitCommonEvents(ctx, msg.Sender)
	return &sdk.Result{Events: ctx.EventManager().Events(), Log: "User successfully invited"}, nil
}

func handleMsgDepositIntoPool(ctx sdk.Context, keeper Keeper, msg types.MsgDepositIntoLiquidityPool) (*sdk.Result, error) {
	if err := keeper.DepositIntoPool(ctx, msg.Depositor, msg.Amount); err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, err.Error())
	}

	return &sdk.Result{Events: ctx.EventManager().Events(), Log: "Amount successfully deposited into pool"}, nil
}

func handleMsgAddTrustedSigner(ctx sdk.Context, keeper Keeper, governmentKeeper government.Keeper, msg types.MsgAddTsp) (*sdk.Result, error) {
	if !governmentKeeper.GetGovernmentAddress(ctx).Equals(msg.Government) {
		return nil, sdkErr.Wrap(sdkErr.ErrInvalidAddress, fmt.Sprintf("Invalid government address: %s", msg.Government))
	}

	membership, err := keeper.GetMembership(ctx, msg.Tsp)
	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrInvalidAddress, fmt.Sprintf("Tsp %s has no membership", msg.Tsp))
	}

	if membership.MembershipType != types.MembershipTypeBlack {
		return nil, sdkErr.Wrap(sdkErr.ErrInvalidAddress, fmt.Sprintf("Membership of Tsp %s is %s but must be %s", msg.Tsp, membership.MembershipType, types.MembershipTypeBlack))
	}

	keeper.AddTrustedServiceProvider(ctx, msg.Tsp)
	ctypes.EmitCommonEvents(ctx, msg.Government)
	return &sdk.Result{Events: ctx.EventManager().Events(), Log: "Tsp successfully added"}, nil
}

func handleMsgRemoveTrustedSigner(ctx sdk.Context, keeper Keeper, governmentKeeper government.Keeper, msg types.MsgRemoveTsp) (*sdk.Result, error) {
	if !governmentKeeper.GetGovernmentAddress(ctx).Equals(msg.Government) {
		return nil, sdkErr.Wrap(sdkErr.ErrInvalidAddress, fmt.Sprintf("Invalid government address: %s", msg.Government))
	}

	keeper.RemoveTrustedServiceProvider(ctx, msg.Tsp)
	return &sdk.Result{Events: ctx.EventManager().Events(), Log: "Tsp successfully removed"}, nil
}

// handleMsgBuyMembership allows to handle a MsgBuyMembership message.
// In order to be able to buy a membership the following requirements must be met.
// 1. The user has been invited from a member already having a membership
// 2. The membership must be valid
// 3. The user has enough stable credits in his wallet
func handleMsgBuyMembership(ctx sdk.Context, keeper Keeper, msg types.MsgBuyMembership) (*sdk.Result, error) {
	// 1. Check the invitation and the invitee membership type
	invite, found := keeper.GetInvite(ctx, msg.Buyer)
	if !found {
		return nil, sdkErr.Wrap(sdkErr.ErrUnauthorized, "Cannot buy a membership without being invited")
	}

	if invite.Status == types.InviteStatusInvalid {
		return nil, sdkErr.Wrap(sdkErr.ErrUnauthorized, fmt.Sprintf("invite for account %s has been marked as invalid previously, cannot continue", msg.Buyer))
	}

	// 2. Verify the membership validity
	if !types.IsMembershipTypeValid(msg.MembershipType) {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("Invalid membership type: %s", msg.MembershipType))
	}

	if !keeper.IsTrustedServiceProvider(ctx, msg.Tsp) {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("Invalid trust service provider: %s", msg.Tsp))
	}

	// Compute expiry height
	height := keeper.ComputeExpiryHeight(ctx.BlockHeight())

	// Allow him to buy the membership
	if err := keeper.BuyMembership(ctx, msg.Buyer, msg.MembershipType, msg.Tsp, height); err != nil {
		return nil, err
	}

	// Give the reward to the invitee
	if err := keeper.DistributeReward(ctx, invite); err != nil {
		return nil, err
	}
	ctypes.EmitCommonEvents(ctx, msg.Tsp)
	return &sdk.Result{Events: ctx.EventManager().Events(), Log: "Membership successfully purchased"}, nil
}

// handleMsgRemoveMembership allows to handle a MsgRemoveMembership message.
// It checks that whoever sent the message is actually the government and remove membership
func handleMsgRemoveMembership(ctx sdk.Context, keeper Keeper, msg types.MsgRemoveMembership) (*sdk.Result, error) {
	govAddr := keeper.governmentKeeper.GetGovernmentAddress(ctx)
	if !govAddr.Equals(msg.Government) {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownAddress,
			fmt.Sprintf("%s is not a government address", msg.Government.String()),
		)
	}
	// TODO add control for tsp?
	err := keeper.RemoveMembership(ctx, msg.Subscriber)
	ctypes.EmitCommonEvents(ctx, msg.Government)
	return &sdk.Result{Events: ctx.EventManager().Events(), Log: "Membership successfully removed"}, err
}

// handleMsgSetMembership handles MsgSetMembership messages.
// It checks that whoever sent the message is actually the government, assigns the membership and then
// distribute the reward to the inviter.
// If the user isn't invited already, an invite will be created.
func handleMsgSetMembership(ctx sdk.Context, keeper Keeper, msg types.MsgSetMembership) (*sdk.Result, error) {
	govAddr := keeper.governmentKeeper.GetGovernmentAddress(ctx)
	if !govAddr.Equals(msg.Government) {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownAddress,
			fmt.Sprintf("%s is not a government address", msg.Government.String()),
		)
	}

	if !types.IsMembershipTypeValid(msg.NewMembership) {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("invalid membership type: %s", msg.NewMembership))
	}

	invite, err := governmentInvitesUser(ctx, keeper, msg.Subscriber)
	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnauthorized, "government could not invite user")
	}

	height := keeper.ComputeExpiryHeight(ctx.BlockHeight())

	err = keeper.AssignMembership(ctx, msg.Subscriber, msg.NewMembership, govAddr, height)
	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest,
			fmt.Sprintf("could not assign membership to user %s: %s", msg.Subscriber, err.Error()),
		)
	}

	err = keeper.DistributeReward(ctx, invite)
	if err != nil {
		return nil, sdkErr.Wrap(sdkErr.ErrUnknownRequest,
			fmt.Sprintf("could not distribute membership reward to user %s: %s", invite.Sender, err.Error()),
		)
	}
	ctypes.EmitCommonEvents(ctx, msg.Government)
	return &sdk.Result{Events: ctx.EventManager().Events(), Log: "Membership successfully set up"}, nil
}

// governmentInvitesUser makes government invite an user if it isn't already invited and validated.
// This function is used when there's the need to assign an arbitrary membership to a given user.
func governmentInvitesUser(ctx sdk.Context, keeper Keeper, user sdk.AccAddress) (types.Invite, error) {
	govAddr := keeper.governmentKeeper.GetGovernmentAddress(ctx)

	// check the user has already been invited
	// if there's an invite, save a credential for it,
	// this way invited, but non-verified users will be able to receive a membership
	invite, found := keeper.GetInvite(ctx, user)
	if found {
		return invite, nil
	}

	// otherwise, create an invite from the government
	err := keeper.InviteUser(ctx, user, govAddr)
	if err != nil {
		return types.Invite{}, err
	}

	// get the invite again, mark it as rewarded, and return it
	invite, found = keeper.GetInvite(ctx, user)
	if !found {
		return types.Invite{}, fmt.Errorf("invite from government created correctly, but invite lookup failed")
	}

	invite.Status = types.InviteStatusRewarded
	keeper.SaveInvite(ctx, invite)

	return invite, nil
}
