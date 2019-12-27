package keeper

import (
	"fmt"
	"strings"

	"github.com/commercionetwork/commercionetwork/x/memberships/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

var membershipCosts = map[string]int64{
	types.MembershipTypeBronze: 25,
	types.MembershipTypeSilver: 250,
	types.MembershipTypeGold:   2500,
	types.MembershipTypeBlack:  25000,
}

type Keeper struct {
	Cdc          *codec.Codec
	StoreKey     sdk.StoreKey
	SupplyKeeper supply.Keeper
}

// NewKeeper creates new instances of the accreditation module Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, supplyKeeper supply.Keeper) Keeper {

	// ensure commerciomint module account is set
	if addr := supplyKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	return Keeper{
		Cdc:          cdc,
		StoreKey:     storeKey,
		SupplyKeeper: supplyKeeper,
	}
}

// storageForAddr returns a string representing the KVStore storage key for an addr.
func storageForAddr(addr sdk.AccAddress) []byte {
	return []byte(types.MembershipsStorageKey + addr.String())
}

// BuyMembership allow to commerciomint and assign a membership of the given membershipType to the specified user.
// If the user already has a membership assigned, deletes the current one and assigns to it the new one.
func (k Keeper) BuyMembership(ctx sdk.Context, buyer sdk.AccAddress, membershipType string) sdk.Error {
	// Get the tokens from the buyer account
	membershipPrice := membershipCosts[membershipType] * 1000000 // Always multiply by one million
	membershipCost := sdk.NewCoins(sdk.NewInt64Coin(k.GetStableCreditsDenom(ctx), membershipPrice))
	if err := k.SupplyKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, membershipCost); err != nil {
		return err
	}
	if err := k.SupplyKeeper.BurnCoins(ctx, types.ModuleName, membershipCost); err != nil {
		return err
	}

	// Assign the membership
	if err := k.AssignMembership(ctx, buyer, membershipType); err != nil {
		return err
	}

	return nil
}

// AssignMembership allow to commerciomint and assign a membership of the given membershipType to the specified user.
// If the user already has a membership assigned, deletes the current one and assigns to it the new one.
// Returns the URI of the new minted token represented the assigned membership, or an error if something goes w
func (k Keeper) AssignMembership(ctx sdk.Context, user sdk.AccAddress, membershipType string) sdk.Error {
	// Check the membership type validity
	if !types.IsMembershipTypeValid(membershipType) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid membership type: %s", membershipType))
	}

	_ = k.RemoveMembership(ctx, user)

	store := ctx.KVStore(k.StoreKey)

	staddr := storageForAddr(user)
	if store.Has(staddr) {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf(
				"cannot add membership \"%s\" for address %s: user already have a membership",
				membershipType,
				user.String(),
			),
		)
	}

	store.Set(staddr, k.Cdc.MustMarshalBinaryBare(membershipType))

	return nil
}

// GetMembership allows to retrieve any existent membership for the specified user.
// The second returned false (the boolean one) tells if the NFT token representing the membership was found or not
func (k Keeper) GetMembership(ctx sdk.Context, user sdk.AccAddress) (types.Membership, sdk.Error) {
	store := ctx.KVStore(k.StoreKey)

	if !store.Has(storageForAddr(user)) {
		return types.Membership{}, sdk.ErrUnknownRequest(
			fmt.Sprintf("membership not found for user \"%s\"", user.String()),
		)
	}

	membershipRaw := store.Get(storageForAddr(user))
	var ms types.Membership
	k.Cdc.MustUnmarshalBinaryBare(membershipRaw, &ms.MembershipType)
	ms.Owner = user

	return ms, nil
}

// RemoveMembership allows to remove any existing membership associated with the given user.
func (k Keeper) RemoveMembership(ctx sdk.Context, user sdk.AccAddress) sdk.Error {
	store := ctx.KVStore(k.StoreKey)

	if !store.Has(storageForAddr(user)) {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("account \"%s\" does not have any membership", user.String()),
		)
	}

	store.Delete(storageForAddr(user))

	return nil
}

// Get GetMembershipsSet returns the list of all the memberships
// that have been minted and are currently stored inside the store
func (k Keeper) GetMembershipsSet(ctx sdk.Context) (types.Memberships, error) {
	store := ctx.KVStore(k.StoreKey)

	ms := types.Memberships{}

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.MembershipsStorageKey))
	for ; iterator.Valid(); iterator.Next() {
		rawAddress := strings.TrimPrefix(string(iterator.Key()), "accreditations:storage:")
		address, err := sdk.AccAddressFromBech32(rawAddress)
		if err != nil {
			return types.Memberships{}, sdk.ErrUnknownRequest(
				fmt.Sprintf("found membership with invalid address: \"%s\"", rawAddress),
			)
		}
		membership := ""
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &membership)
		ms = append(ms, types.Membership{
			MembershipType: membership,
			Owner:          address,
		})
	}

	return ms, nil
}

// GetStableCreditsDenom returns the denom that must be used when referring to stable credits
// that can be used to purchase a membership
func (k Keeper) GetStableCreditsDenom(ctx sdk.Context) (denom string) {
	store := ctx.KVStore(k.StoreKey)
	return string(store.Get([]byte(types.StableCreditsStoreKey)))
}

// SetStableCreditsDenom allows to set the denom of the coins that must be used as stable credits
// when purchasing a membership.
func (k Keeper) SetStableCreditsDenom(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.StableCreditsStoreKey), []byte(denom))
}
