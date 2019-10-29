package keeper

import (
	"fmt"
	"strings"

	"github.com/commercionetwork/commercionetwork/x/memberships/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/cosmos/cosmos-sdk/x/nft/exported"
)

var membershipCosts = map[string]int64{
	types.MembershipTypeBronze: 25,
	types.MembershipTypeSilver: 250,
	types.MembershipTypeGold:   2500,
	types.MembershipTypeBlack:  25000,
}

type Keeper struct {
	cdc        *codec.Codec
	StoreKey   sdk.StoreKey
	BankKeeper bank.Keeper
	NftKeeper  nft.Keeper
}

// NewKeeper creates new instances of the accreditation module Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, nftK nft.Keeper, bankK bank.Keeper) Keeper {
	return Keeper{
		StoreKey:   storeKey,
		BankKeeper: bankK,
		NftKeeper:  nftK,
		cdc:        cdc,
	}
}

// getMembershipTokenId allows to retrieve the id of a token representing a membership associated to the given user
func (keeper Keeper) getMembershipTokenId(user sdk.AccAddress) string {
	return "membership-" + user.String()
}

// getMembershipUri allows to returns the URI of the NFT that represents a membership of the
// given membershipType and having the given id
func (keeper Keeper) getMembershipUri(membershipType string, id string) string {
	return fmt.Sprintf("membership:%s:%s", membershipType, id)
}

// BuyMembership allow to mint and assign a membership of the given membershipType to the specified user.
// If the user already has a membership assigned, deletes the current one and assigns to it the new one.
func (keeper Keeper) BuyMembership(ctx sdk.Context, buyer sdk.AccAddress, membershipType string) sdk.Error {
	// Get the tokens from the buyer account
	membershipPrice := membershipCosts[membershipType] * 1000000 // Always multiply by one million
	membershipCost := sdk.NewCoins(sdk.NewInt64Coin(keeper.GetStableCreditsDenom(ctx), membershipPrice))
	if _, err := keeper.BankKeeper.SubtractCoins(ctx, buyer, membershipCost); err != nil {
		return err
	}

	if _, err := keeper.AssignMembership(ctx, buyer, membershipType); err != nil {
		return err
	}

	return nil
}

// AssignMembership allow to mint and assign a membership of the given membershipType to the specified user.
// If the user already has a membership assigned, deletes the current one and assigns to it the new one.
// Returns the URI of the new minted token represented the assigned membership, or an error if something goes w
func (keeper Keeper) AssignMembership(ctx sdk.Context, user sdk.AccAddress, membershipType string) (string, sdk.Error) {
	// Check the membership type validity
	if !types.IsMembershipTypeValid(membershipType) {
		return "", sdk.ErrUnknownRequest("Invalid membership type")
	}

	// Find any existing membership
	if _, err := keeper.RemoveMembership(ctx, user); err != nil {
		return "", sdk.ErrUnknownRequest(err.Error())
	}

	// Build the token information
	id := keeper.getMembershipTokenId(user)
	uri := keeper.getMembershipUri(membershipType, id)

	// Build the membership token
	membershipToken := nft.NewBaseNFT(id, user, uri)

	// Mint the token
	if err := keeper.NftKeeper.MintNFT(ctx, types.NftDenom, &membershipToken); err != nil {
		return "", err
	}

	// Return with no error
	return membershipToken.TokenURI, nil
}

// GetMembership allows to retrieve any existent membership for the specified user.
// The second returned false (the boolean one) tells if the NFT token representing the membership was found or not
func (keeper Keeper) GetMembership(ctx sdk.Context, user sdk.AccAddress) (exported.NFT, bool) {
	foundToken, err := keeper.NftKeeper.GetNFT(ctx, types.NftDenom, keeper.getMembershipTokenId(user))

	// The token was not found
	if err != nil {
		return nil, false
	}

	return foundToken, true
}

// RemoveMembership allows to remove any existing membership associated with the given user.
func (keeper Keeper) RemoveMembership(ctx sdk.Context, user sdk.AccAddress) (bool, sdk.Error) {
	id := keeper.getMembershipTokenId(user)

	if found, _ := keeper.NftKeeper.GetNFT(ctx, types.NftDenom, id); found == nil {
		// The token was not found, so it's trivial to delete it: simply do nothing
		return true, nil
	}

	if err := keeper.NftKeeper.DeleteNFT(ctx, types.NftDenom, keeper.getMembershipTokenId(user)); err != nil {
		// The token was found, but an error was raised during the deletion. Return the error
		return false, err
	}

	// The token was found and deleted
	return true, nil
}

// GetMembershipType returns the type of the membership represented by the given NFT token
func (keeper Keeper) GetMembershipType(membership exported.NFT) string {
	return strings.Split(membership.GetTokenURI(), ":")[1]
}

// Get GetMembershipsSet returns the list of all the memberships
// that have been minted and are currently stored inside the store
func (keeper Keeper) GetMembershipsSet(ctx sdk.Context) []types.Membership {

	collection, found := keeper.NftKeeper.GetCollection(ctx, types.NftDenom)
	if !found {
		return nil
	}

	memberships := make([]types.Membership, len(collection.NFTs))
	for index, membershipNft := range collection.NFTs {
		memberships[index] = types.Membership{
			Owner:          membershipNft.GetOwner(),
			MembershipType: keeper.GetMembershipType(membershipNft),
		}
	}

	return memberships
}

// GetStableCreditsDenom returns the denom that must be used when referring to stable credits
// that can be used to purchase a membership
func (keeper Keeper) GetStableCreditsDenom(ctx sdk.Context) (denom string) {
	store := ctx.KVStore(keeper.StoreKey)
	keeper.cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.StableCreditsStoreKey)), &denom)
	return denom
}

// SetStableCreditsDenom allows to set the denom of the coins that must be used as stable credits
// when purchasing a membership.
func (keeper Keeper) SetStableCreditsDenom(ctx sdk.Context, denom string) {
	store := ctx.KVStore(keeper.StoreKey)
	store.Set([]byte(types.StableCreditsStoreKey), keeper.cdc.MustMarshalBinaryBare(&denom))
}
