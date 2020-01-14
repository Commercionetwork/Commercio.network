package keeper

import (
	"fmt"

	"github.com/commercionetwork/commercionetwork/x/memberships/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		// get all the users with membership
		i := k.MembershipIterator(ctx)
		defer i.Close()
		for ; i.Valid(); i.Next() {
			user := k.ExtractMembership(i.Key(), i.Value())
			credentials := k.GetUserCredentials(ctx, user.Owner)

			// check that the user has been invited
			_, found := k.GetInvite(ctx, user.Owner)
			if !found {
				return sdk.FormatInvariant(
					types.ModuleName,
					membershipVerifiedInvName,
					fmt.Sprintf(
						"found user with membership but with no invite: %s",
						user.Owner.String(),
					),
				), true
			}

			// check that there are credentials for user
			if len(credentials) == 0 {
				return sdk.FormatInvariant(
					types.ModuleName,
					membershipVerifiedInvName,
					fmt.Sprintf(
						"found user with membership but with no credentials: %s",
						user.Owner.String(),
					),
				), true
			}

			// for each credential, check that the Verifier is actually
			// a tsp
			for _, credential := range credentials {
				if !k.IsTrustedServiceProvider(ctx, credential.Verifier) {
					return sdk.FormatInvariant(
						types.ModuleName,
						membershipVerifiedInvName,
						fmt.Sprintf(
							"found user whose credential was verified by a non-Verifier %s user but with no credentials: %s",
							credential.Verifier.String(),
							user.Owner.String(),
						),
					), true
				}
			}
		}

		return "", false
	}
}
