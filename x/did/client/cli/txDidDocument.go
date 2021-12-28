package cli

import (
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/commercionetwork/commercionetwork/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdSetDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-DID-document [did_document_proposal_path]",
		Short: "Sets the DID document for the requesting address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			//keybase, err := keyring.New(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), os.Stdin)

			argsDDOpath, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			// read document from path
			ddoData, err := ioutil.ReadFile(argsDDOpath)
			if err != nil {
				return err
			}

			var didDocumentProposal types.MsgSetDidDocument
			json.Unmarshal(ddoData, &didDocumentProposal)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if err := didDocumentProposal.ValidateBasic(); err != nil {
				return err
			}

			// check requesting address equals to DID subject

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &didDocumentProposal)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

/*
func signDidDocument(cliCtx client.Context, unsignedDoc types.DidDocument, keybase keyring.Keyring) (string, error) {
	cdc := codec.NewLegacyAmino()
	jsonUnsigned, err := cdc.MarshalJSON(unsignedDoc)
	if err != nil {
		return "", fmt.Errorf("error marshaling doc into json")
	}
	_ = jsonUnsigned
	sign, _, err := keybase.SignByAddress(cliCtx.GetFromAddress(), jsonUnsigned)
	if err != nil {
		return "", fmt.Errorf("failed to sign tx")
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}
*/

// Calculate Proof
/*signature, err := signDidDocument(clientCtx, didDocument, keybase)

fromAddressPubkey, err := keybase.KeyByAddress(clientCtx.GetFromAddress())
if err != nil {
	return fmt.Errorf("could not get keybase for address: %w", err)
}

verMeth, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, fromAddressPubkey.GetPubKey())
if err != nil {
	return fmt.Errorf("could not derive address public key: %w", err)
}

proof := types.Proof{
	Type:               types.KeyTypeSecp256k12019,
	Created:            time.Now().String(),
	ProofPurpose:       types.ProofPurposeAuthentication,
	Controller:         clientCtx.GetFromAddress().String(),
	VerificationMethod: verMeth,
	SignatureValue:     signature,
}*/
