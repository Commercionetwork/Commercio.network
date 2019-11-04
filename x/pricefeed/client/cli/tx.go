package cli

import (
	"fmt"

	"github.com/commercionetwork/commercionetwork/x/pricefeed/internal/types"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Pricefeed transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(GetCmdSetPrice(cdc), GetCmdAddOracle(cdc))

	return txCmd
}

func GetCmdSetPrice(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-price [token-name] [token-price] [expiry]",
		Short: "set price for a given token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			tokenPrice, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return sdk.ErrInternal(err.Error())
			}

			expiry, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return sdk.ErrInternal(fmt.Sprintf("Invalid expiration height, %s", args[2]))
			}

			price := types.NewPrice(args[0], tokenPrice, expiry)

			oracle := cliCtx.GetFromAddress()
			msg := types.NewMsgSetPrice(price, oracle)

			err2 := msg.ValidateBasic()
			if err2 != nil {
				return err2
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdAddOracle cli command for posting prices.
func GetCmdAddOracle(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-oracle [oracle-address]",
		Short: "add a trusted oracle to the oracles' list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			oracle, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			signer := cliCtx.GetFromAddress()
			msg := types.NewMsgAddOracle(signer, oracle)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
