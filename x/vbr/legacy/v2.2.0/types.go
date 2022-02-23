// DONTCOVER
// nolint
package v2_2_0

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "vbr"
)

// ---------------
// --- Genesis
// ---------------
// v2.2.0 vbr genesis state
type GenesisState struct {
	PoolAmount        sdk.DecCoins `json:"pool_amount"`
	RewardRate        sdk.Dec      `json:"reward_rate"`
	AutomaticWithdraw bool         `json:"automatic_withdraw"`
}


type MsgIncrementBlockRewardsPool struct {
	Funder sdk.AccAddress `json:"funder"`
	Amount sdk.Coins      `json:"amount"`
}

type MsgSetRewardRate struct {
	Government sdk.AccAddress `json:"government"`
	RewardRate sdk.Dec        `json:"reward_rate"`
}

type MsgSetAutomaticWithdraw struct {
	Government        sdk.AccAddress `json:"government"`
	AutomaticWithdraw bool           `json:"automatic_withdraw"`
}
