package types

import (
	"encoding/hex"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DidDepositRequest represents the request that is sent from a user when he wants to use
// some amounts from his public identity to later power up his pairwise Did.
// This request will later be read and unencrypted from a central organization that will move the funds.
type DidDepositRequest struct {
	Status        *RequestStatus `json:"status"`         // Type of the request
	Recipient     sdk.AccAddress `json:"recipient"`      // Address that should be funded
	Amount        sdk.Coins      `json:"amount"`         // Amount that should be taken
	Proof         string         `json:"proof"`          // Proof of the deposit, encrypted using an AES-256 key and hex encoded
	EncryptionKey string         `json:"encryption_key"` // AES-256 key encrypted using reader's public key and hex encoded
	FromAddress   sdk.AccAddress `json:"from_address"`   // Address from which the funds should be taken
}

func ValidateHex(proof string) sdk.Error {

	depositProof := strings.TrimSpace(proof)
	if len(depositProof) == 0 {
		return sdk.ErrUnknownRequest("Hex value cannot be empty")
	}

	if _, err := hex.DecodeString(depositProof); err != nil {
		return sdk.ErrUnknownRequest("Hex value must be hex encoded")
	}

	return nil
}

func ValidateEncryptionKey(key string) sdk.Error {
	encryptionKey := strings.TrimSpace(key)
	if len(encryptionKey) == 0 {
		return sdk.ErrUnknownRequest("Encryption key cannot be empty")
	}

	if _, err := hex.DecodeString(encryptionKey); err != nil {
		return sdk.ErrUnknownRequest("Encryption key must be hex encoded")
	}

	return nil
}
