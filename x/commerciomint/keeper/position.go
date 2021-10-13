package keeper

import (
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"
)

func (k Keeper) SetPosition(ctx sdk.Context, position types.Position) error {
	store := ctx.KVStore(k.storeKey)
	key := makePositionKey(sdk.AccAddress(position.Owner), position.ID)

	if store.Has(key) {
		return fmt.Errorf("a position with id %s already exists", position.ID)
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&position))

	return nil
}

func (k Keeper) UpdatePosition(ctx sdk.Context, position types.Position) error {
	store := ctx.KVStore(k.storeKey)
	key := makePositionKey(sdk.AccAddress(position.Owner), position.ID)

	if !store.Has(key) {
		return fmt.Errorf("a position with id %s doesn't exists", position.ID)
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&position))

	return nil
}

func (k Keeper) GetPosition(ctx sdk.Context, owner sdk.AccAddress, id string) (types.Position, bool) {
	position := types.Position{}
	key := makePositionKey(owner, id)
	store := ctx.KVStore(k.storeKey)
	bs := store.Get(key)
	if bs == nil {
		return position, false
	}
	k.cdc.MustUnmarshalBinaryBare(bs, &position)
	return position, true
}

func (k Keeper) GetAllPositionsOwnedBy(ctx sdk.Context, owner sdk.AccAddress) []*types.Position {
	var positions []*types.Position
	i := k.newPositionsByOwnerIterator(ctx, owner)
	defer i.Close()
	for ; i.Valid(); i.Next() {
		var position types.Position
		k.cdc.MustUnmarshalBinaryBare(i.Value(), &position)
		positions = append(positions, &position)
	}
	return positions
}

// NewPosition creates a new minting position for the amount deposited, credited to depositor.
//func (k Keeper) NewPosition(ctx sdk.Context, depositor sdk.AccAddress, deposit sdk.Coins, id string) error {
func (k Keeper) NewPosition(ctx sdk.Context, position types.Position) error {
	depositor := sdk.AccAddress(position.Owner)
	ucccRequested := sdk.NewInt(position.Collateral)
	if ucccRequested.IsZero() {
		//return errors.New("no uccc requested")
		return fmt.Errorf("no %s requested", types.CreditsDenom)

	}

	conversionRate := k.GetConversionRate(ctx)

	uccDec := sdk.NewDecFromInt(ucccRequested)
	ucommercioAmount := uccDec.Mul(conversionRate.Dec).Ceil().TruncateInt()

	// Create ucccEmitted token
	ucccEmitted := sdk.NewCoin(types.CreditsDenom, ucccRequested)

	ucomAmount := sdk.NewCoin("ucommercio", ucommercioAmount)

	// Create the ETP and validate it
	/*position := types.NewPosition(
		depositor,
		ucomAmount.Amount,
		ucccEmitted,
		id,
		ctx.BlockTime(),
		conversionRate,
	)*/
	position.CreatedAt = ctx.BlockTime().String()
	position.ExchangeRate = &conversionRate
	position.Credits = &ucccEmitted

	if err := position.Validate(); err != nil {
		return fmt.Errorf("could not validate position message, %w", err)
	}

	// Send the deposit from the user to the commerciomint account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoins(ucomAmount)); err != nil {
		return fmt.Errorf("could not move collateral amount to module account, %w", err)
	}

	// Mint the tokens and send them to the user
	creditsCoins := sdk.NewCoins(ucccEmitted)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, creditsCoins); err != nil {
		return fmt.Errorf("could not mint coins, %w", err)
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, creditsCoins); err != nil {
		return fmt.Errorf("could not send minted coins to account, %w", err)
	}

	// Create position
	k.SetPosition(ctx, position)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		eventNewPosition,
		sdk.NewAttribute("depositor", depositor.String()),
		sdk.NewAttribute("amount_deposited", ucomAmount.String()),
		sdk.NewAttribute("minted_coins", creditsCoins.String()),
		sdk.NewAttribute("position_id", position.ID),
		sdk.NewAttribute("timestamp", position.CreatedAt),
	))

	return nil
}

func (k Keeper) GetAllPositions(ctx sdk.Context) []*types.Position {
	var positions []*types.Position
	iterator := k.newPositionsIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var pos types.Position
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &pos)
		positions = append(positions, &pos)
	}

	return positions
}

// BurnCCC burns burnAmount to the conversion rate stored in the Position identified by id, and returns the
// resulting collateral amount to user.
func (k Keeper) RemoveCCC(ctx sdk.Context, user sdk.AccAddress, id string, burnAmount sdk.Coin) (sdk.Int, error) {
	pos, found := k.GetPosition(ctx, user, id)
	residualAmount := sdk.NewInt(0)
	if !found {
		msg := fmt.Sprintf("position for user with address %s and id %s does not exist", user, id)
		return residualAmount, sdkErr.Wrap(sdkErr.ErrUnknownRequest, msg)
	}

	// Control if position is almost in freezing period
	freezePeriod := k.GetFreezePeriod(ctx)
	createdAt, _ := time.Parse(time.RFC3339, pos.CreatedAt) // TODO CHECK FORMAT AND ERROR
	if time.Now().Sub(createdAt) <= freezePeriod {
		return residualAmount, sdkErr.Wrap(sdkErr.ErrInvalidRequest, "cannot burn position yet in the freeze period")
	}

	// Control if tokens request to burn are more then initially requested
	if pos.Credits.Amount.Sub(burnAmount.Amount).IsNegative() {
		return residualAmount, sdkErr.Wrap(sdkErr.ErrInvalidRequest, "cannot burn more tokens that those initially requested")
	}

	shouldDeletePos := pos.Credits.Amount.Sub(burnAmount.Amount).IsZero()

	// 1. burn burnAmount tokens
	// 2. decrement pos.Credits
	// 3. give user amounts of collateral back
	// 4. decrement collateral
	// 5. save or delete position

	// 1.
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, sdk.NewCoins(burnAmount))
	if err != nil {
		return residualAmount, sdkErr.Wrapf(sdkErr.ErrInvalidRequest, "cannot send tokens from sender to module, %s", err.Error())
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burnAmount))
	if err != nil {
		return residualAmount, sdkErr.Wrapf(sdkErr.ErrInvalidRequest, "cannot burn coins, %s", err)
	}

	// 2.
	residualAmount = pos.Credits.Amount.Sub(burnAmount.Amount)
	*pos.Credits = sdk.NewCoin(
		pos.Credits.Denom,
		residualAmount,
	)

	// 3.
	burnAmountDec := sdk.NewDecFromInt(burnAmount.Amount)
	collateralAmount := burnAmountDec.Mul(pos.ExchangeRate.Dec).Ceil().TruncateInt()
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, sdk.NewCoins(sdk.NewCoin(
		types.BondDenom,
		collateralAmount,
	)))
	if err != nil {
		return residualAmount, sdkErr.Wrapf(sdkErr.ErrInvalidRequest, "cannot send collateral from module to sender, %s", err.Error())
	}

	// 4.
	posCollateral := sdk.NewInt(pos.Collateral)
	diffCollateral := posCollateral.Sub(collateralAmount) // TODO CONVERT INT64 TO COIN
	pos.Collateral = diffCollateral.Int64()               // TODO Should panic

	// TODO Review events
	defer func(deleted bool, ctx sdk.Context) {
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			eventBurnCCC,
			sdk.NewAttribute("position_id", pos.ID),
			sdk.NewAttribute("sender", user.String()),
			sdk.NewAttribute("amount", burnAmount.String()),
			sdk.NewAttribute("position_deleted", strconv.FormatBool(shouldDeletePos))))
	}(shouldDeletePos, ctx)

	// 5.
	if shouldDeletePos {
		residualAmount = sdk.NewInt(0)
		k.deletePosition(ctx, pos)
		return residualAmount, nil
	}

	k.UpdatePosition(ctx, pos)

	return residualAmount, nil
}

func (k Keeper) newPositionsByOwnerIterator(ctx sdk.Context, owner sdk.AccAddress) sdk.Iterator {
	prefix := append([]byte(types.EtpStorePrefix), owner...)
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), prefix)
}

func makePositionKey(address sdk.AccAddress, id string) []byte {
	base := append([]byte(types.EtpStorePrefix), address...)
	return append(base, []byte(id)...)
}

func (k Keeper) newPositionsIterator(ctx sdk.Context) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), []byte(types.EtpStorePrefix))
}

func (k Keeper) deletePosition(ctx sdk.Context, pos types.Position) {
	store := ctx.KVStore(k.storeKey)
	key := makePositionKey(sdk.AccAddress(pos.Owner), pos.ID)
	if bs := store.Get(key); bs == nil {
		panic(fmt.Sprintf("no pos stored at key %s", key))
	}
	store.Delete(key)
}