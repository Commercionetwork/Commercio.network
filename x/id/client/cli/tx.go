package cli

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/commercionetwork/commercionetwork/x/id/types"
)

const flagPrivRsaVerKey = "privRsaVerKey"
const flagPrivRsaSignKey = "privRsaSignKey"

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "CommercioDOCS id subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		getSetIdentityCommand(cdc),
	)

	return txCmd
}

func getSetIdentityCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setidentity",
		Short: "sets the identity",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			keybase, err := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), os.Stdin)
			if err != nil {
				return fmt.Errorf("error accesing to keyring: %w", err)
			}

			verPubKey, err := getVerificationPublicKey(cliCtx, viper.GetString(flagPrivRsaVerKey))
			if err != nil {
				return fmt.Errorf("error extracting public from private key verification path: %s", err)
			}

			signPubKey, err := getSignPublicKey(cliCtx, viper.GetString(flagPrivRsaSignKey))
			if err != nil {
				return fmt.Errorf("error extracting public from private key sign path: %s", err)
			}

			unsignedDoc := types.DidDocumentUnsigned{
				Context: types.ContextDidV1,
				ID:      cliCtx.GetFromAddress(),
				PubKeys: types.PubKeys{
					verPubKey,
					signPubKey,
				},
			}

			signature, err := signDidDocument(cliCtx, cdc, unsignedDoc, keybase)
			if err != nil {
				return err
			}

			fromAddressPubkey, err := keybase.GetByAddress(cliCtx.GetFromAddress())
			if err != nil {
				return fmt.Errorf("could not get keybase for address: %w", err)
			}

			verMeth, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, fromAddressPubkey.GetPubKey())
			if err != nil {
				return fmt.Errorf("could not derive address public key: %w", err)
			}

			proof := types.Proof{
				Type:               types.KeyTypeSecp256k12019,
				Created:            time.Now(),
				ProofPurpose:       types.ProofPurposeAuthentication,
				Controller:         cliCtx.GetFromAddress().String(),
				VerificationMethod: verMeth,
				SignatureValue:     signature,
			}

			msg := types.NewMsgSetIdentity(types.DidDocument{
				Context: unsignedDoc.Context,
				ID:      unsignedDoc.ID,
				PubKeys: unsignedDoc.PubKeys,
				Proof:   proof,
				Service: nil,
			})

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

		},
	}

	cmd.Flags().String(flagPrivRsaSignKey, "", "the path of the pem rsa sign key")
	cmd.MarkFlagRequired(flagPrivRsaSignKey)

	cmd.Flags().String(flagPrivRsaVerKey, "", "the path of the pem rsa verification key")
	cmd.MarkFlagRequired(flagPrivRsaVerKey)

	cmd = flags.PostCommands(cmd)[0]
	cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func signDidDocument(cliCtx context.CLIContext, cdc *codec.Codec, unsignedDoc types.DidDocumentUnsigned, keybase keys.Keybase) (string, error) {
	jsonUnsigned, err := cdc.MarshalJSON(unsignedDoc)
	if err != nil {
		return "", fmt.Errorf("error marshaling doc into json")
	}

	sign, _, err := keybase.Sign(cliCtx.GetFromName(), "", jsonUnsigned)
	if err != nil {
		return "", fmt.Errorf("failed to sign tx")
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func getVerificationPublicKey(cliCtx context.CLIContext, path string) (types.PubKey, error) {
	pk, err := types.LoadRSAPrivKeyFromDisk(path)
	if err != nil {
		return types.PubKey{}, err
	}

	fromAddress := cliCtx.GetFromAddress()
	pemString, err := types.PublicKeyToPemString(&pk.PublicKey)
	if err != nil {
		return types.PubKey{}, err
	}

	verPubKey := types.PubKey{
		ID:         fmt.Sprintf("%s#keys-1", fromAddress.String()),
		Type:       types.KeyTypeRsaVerification,
		Controller: fromAddress,
		PublicKey:  pemString,
	}

	return verPubKey, nil
}

func getSignPublicKey(cliCtx context.CLIContext, path string) (types.PubKey, error) {
	pk, err := types.LoadRSAPrivKeyFromDisk(path)
	if err != nil {
		return types.PubKey{}, err
	}

	fromAddress := cliCtx.GetFromAddress()
	pemString, err := types.PublicKeyToPemString(&pk.PublicKey)
	if err != nil {
		return types.PubKey{}, err
	}

	verPubKey := types.PubKey{
		ID:         fmt.Sprintf("%s#keys-2", fromAddress.String()),
		Type:       types.KeyTypeRsaSignature,
		Controller: fromAddress,
		PublicKey:  pemString,
	}

	return verPubKey, nil
}
