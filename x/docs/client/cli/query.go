package cli

import (
	"fmt"

	"github.com/commercionetwork/commercionetwork/x/docs/internal/types"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetCmdSentDocuments(cdc), GetCmdReceivedDocuments(cdc), GetCmdReceivedReceipts(cdc))

	return cmd
}

// ----------------------------------
// --- ShareDocument
// ----------------------------------

func GetCmdRetrieveDocument(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "document [document-uuid]",
		Short: "Get the document information for the given uuid value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			checksumValue := args[0]

			route := fmt.Sprintf("custom/%s/document/%s", types.QuerierRoute, checksumValue)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not get document from checksum %s: \n %s", checksumValue, err)
				return nil
			}

			fmt.Println(string(res))

			return nil
		},
	}
}

func GetCmdSentDocuments(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sent-documents [user-address]",
		Short: "Get all documents sent by user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QuerySentDocuments, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not get sent documents by user: \n %s", err)
			}

			fmt.Println(string(res))

			return nil
		},
	}
}

func GetCmdReceivedDocuments(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "received-documents [user-address]",
		Short: "Get all documents received by user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryReceivedDocuments, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not get received documents by user: \n %s", err)
			}

			fmt.Println(string(res))

			return nil
		},
	}
}

func GetCmdSharedDocumentsWithUser(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "shared-with [user-address]",
		Short: "Get all documents shared with given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				fmt.Printf("The given address doesn't exist or it's wrong")
			}

			route := fmt.Sprintf("custom/%s/shared-with/%s", types.QuerierRoute, address)
			res, _, err2 := cliCtx.QueryWithData(route, nil)
			if err2 != nil {
				fmt.Printf("Could not get any shared document with %s: \n %s", string(address), err2)
			}

			fmt.Println(string(res))

			return nil
		},
	}
}

// ----------------------------------
// --- DocumentReceipt
// ----------------------------------

func GetCmdReceivedReceipts(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "receipts [user-address] [[doc-uuid]]",
		Short: "Get the document receipt associated with given document uuid",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			addr, uuid := args[0], ""
			if len(args) == 2 {
				uuid = args[1]
			}

			route := fmt.Sprintf("custom/%s/%s/%s/%s", types.QuerierRoute, types.QueryReceipts, addr, uuid)
			res, _, err2 := cliCtx.QueryWithData(route, nil)
			if err2 != nil {
				fmt.Printf("Could not get any receipt associated with the given user or uuid: \n %s", err2)
			}

			fmt.Printf(string(res))

			return nil
		},
	}
}
