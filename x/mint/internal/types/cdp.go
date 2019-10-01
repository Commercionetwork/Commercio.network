package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//CDP stands for Collateralized Debt Position
type CDP struct {
	Owner           sdk.AccAddress `json:"owner"`
	DepositedAmount sdk.Coins      `json:"deposited_amount"`
	LiquidityAmount sdk.Coins      `json:"liquidity_amount"`
	Timestamp       string         `json:"timestamp"`
}

func NewCDP(request CDPRequest, liquidityAmount sdk.Coins) CDP {
	return CDP{
		Owner:           request.Signer,
		DepositedAmount: request.DepositedAmount,
		LiquidityAmount: liquidityAmount,
		Timestamp:       request.Timestamp,
	}
}

func (current CDP) Equals(cdp CDP) bool {
	return current.Owner.Equals(cdp.Owner) &&
		current.DepositedAmount.IsEqual(cdp.DepositedAmount) &&
		current.LiquidityAmount.IsEqual(cdp.LiquidityAmount) &&
		current.Timestamp == cdp.Timestamp
}

type CDPs []CDP

func (cdps CDPs) AppendIfMissing(cdp CDP) (CDPs, bool) {
	for _, ele := range cdps {
		if ele.Equals(cdp) {
			return nil, true
		}
	}
	return append(cdps, cdp), false
}

//This method filters a slice without allocating a new underlying array
func (cdps CDPs) RemoveWhenFound(timestamp string) (CDPs, bool) {
	tmp := cdps[:0]
	found := false
	for _, ele := range cdps {
		if ele.Timestamp != timestamp {
			tmp = append(tmp, ele)
		} else {
			found = true
		}
	}
	return tmp, found
}

func (cdps CDPs) GetCdpFromTimestamp(timestamp string) (*CDP, bool) {
	for _, ele := range cdps {
		if ele.Timestamp == timestamp {
			return &ele, true
		}
	}
	return nil, false
}
