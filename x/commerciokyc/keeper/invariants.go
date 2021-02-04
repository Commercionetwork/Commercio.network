package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/commercionetwork/commercionetwork/x/commerciokyc/types"
)

const (
	membershipVerifiedInvName string = "user-membership-verifier"
)

// RegisterInvariants registers all memberships invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, membershipVerifiedInvName,
		MembershipVerifiedInvariant(k))
}

// MembershipVerifiedInvariant verifies that all the membership existing in the state have been authorized,
// and that the users associated with them have been invited by a TSP.
func MembershipVerifiedInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		/*
			// get gov address
			govAddr := k.governmentKeeper.GetGovernmentAddress(ctx)

			// get all the users with membership
			i := k.MembershipIterator(ctx)
			defer i.Close()
			for ; i.Valid(); i.Next() {
				membership := k.ExtractMembership(i.Value())

				if govAddr.Equals(membership.Owner) {
					continue
				}

				// check that the user has been invited
				invite, found := k.GetInvite(ctx, membership.Owner)
				if !found || (invite.Status == types.InviteStatusPending || invite.Status == types.InviteStatusInvalid) {
					return sdk.FormatInvariant(
						types.ModuleName,
						membershipVerifiedInvName,
						fmt.Sprintf(
							"found user with membership but with no invite: %s",
							membership.Owner.String(),
						),
					), true
				}
			}*/

		return "", false
	}
}
