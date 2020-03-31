package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"
)

const (
	cdpsForExistingPrice       string = "cdp-existing-price"
	liquidityPoolSumEqualsCdps string = "liquidity-pool-sum-equals-cdps"
)

func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, cdpsForExistingPrice,
		CdpsForExistingPrice(k))
	ir.RegisterRoute(types.ModuleName, liquidityPoolSumEqualsCdps,
		LiquidityPoolAmountEqualsCdps(k))
}

// CdpsForExistingPrice checks that each Cdp currently opened refers to an existing token priced by x/pricefeed.
func CdpsForExistingPrice(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		cdps := k.GetCdps(ctx)

		for _, cdp := range cdps {
			for _, depositAmount := range cdp.DepositedAmount {
				price, ok := k.priceFeedKeeper.GetCurrentPrice(ctx, depositAmount.Denom)
				if !ok || price.Value.IsZero() {
					return sdk.FormatInvariant(
						types.ModuleName,
						cdpsForExistingPrice,
						fmt.Sprintf(
							"found cdp from owner %s which refers to a nonexistent asset %s for %s amount",
							cdp.Owner.String(),
							depositAmount.Denom,
							depositAmount.Amount.String(),
						),
					), true
				}
			}
		}
		return "", false
	}
}

// LiquidityPoolAmountEqualsCdps checks that the value of all the opened cdps equals the liquidity pool amount.
func LiquidityPoolAmountEqualsCdps(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		cdps := k.GetCdps(ctx)

		sums := make(map[string]sdk.Int)
		for _, cdp := range cdps {
			for _, depositAmount := range cdp.DepositedAmount {
				if _, ok := sums[depositAmount.Denom]; !ok {
					sums[depositAmount.Denom] = sdk.NewInt(0)
				}

				sums[depositAmount.Denom] = sums[depositAmount.Denom].Add(depositAmount.Amount)
			}
		}

		pool := k.GetLiquidityPoolAmount(ctx)

		if pool.IsZero() && len(cdps) > 0 {
			return sdk.FormatInvariant(
				types.ModuleName,
				cdpsForExistingPrice,
				fmt.Sprintf(
					"cdps opened and liquidity pool is empty",
				),
			), true
		}

		for name, sum := range sums {
			for _, token := range pool {
				if token.Denom == name {
					if !sum.Equal(token.Amount) {
						return sdk.FormatInvariant(
							types.ModuleName,
							cdpsForExistingPrice,
							fmt.Sprintf(
								"pool amount for denom %s doesn't correspond to the sum of all the cdps opened for it, which is %s%s",
								name,
								sum.String(),
								name,
							),
						), true
					}
				}
			}
		}

		return "", false
	}
}
