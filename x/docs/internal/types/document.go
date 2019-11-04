package types

import (
	"fmt"
	"strings"

	"github.com/commercionetwork/commercionetwork/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	uuid "github.com/satori/go.uuid"
)

// Document contains the generic information about a single document which has been sent from a user to another user.
// It contains the information about its content, its associated metadata and the related checksum.
// In order to be valid, a document must have a non-empty and unique UUID and a valid metadata information.
// Both the content and the checksum information are optional.
type Document struct {
	Sender         sdk.AccAddress          `json:"sender"`
	Recipients     types.Addresses         `json:"recipients"`
	UUID           string                  `json:"uuid"`
	Metadata       DocumentMetadata        `json:"metadata"`
	ContentURI     string                  `json:"content_uri"`     // Optional
	Checksum       *DocumentChecksum       `json:"checksum"`        // Optional
	EncryptionData *DocumentEncryptionData `json:"encryption_data"` // Optional
}

// TODO: Test
func (doc Document) Equals(other Document) bool {
	validContent := doc.UUID == other.UUID &&
		doc.ContentURI == other.ContentURI &&
		doc.Metadata.Equals(other.Metadata)

	var validChecksum bool
	if doc.Checksum != nil && other.Checksum != nil {
		validChecksum = doc.Checksum.Equals(*other.Checksum)
	} else {
		validChecksum = doc.Checksum == other.Checksum
	}

	var validEncryptionData bool
	if doc.EncryptionData != nil && other.EncryptionData != nil {
		validEncryptionData = doc.EncryptionData.Equals(*other.EncryptionData)
	} else {
		validEncryptionData = doc.EncryptionData == other.EncryptionData
	}

	return validContent && validChecksum && validEncryptionData
}

func validateUUID(uuidStr string) bool {
	_, err := uuid.FromString(uuidStr)

	// when err is nil, uuidStr is a valid UUID
	return err == nil
}

// TODO: Test
func (doc Document) Validate() sdk.Error {
	if doc.Sender.Empty() {
		return sdk.ErrInvalidAddress(doc.Sender.String())
	}

	if doc.Recipients.Empty() {
		return sdk.ErrInvalidAddress("Recipients cannot be empty")
	}

	for _, recipient := range doc.Recipients {
		if recipient.Empty() {
			return sdk.ErrInvalidAddress(recipient.String())
		}
	}

	if !validateUUID(doc.UUID) {
		return sdk.ErrUnknownRequest("Invalid document UUID")
	}

	err := doc.Metadata.Validate()
	if err != nil {
		return sdk.ErrUnknownRequest(err.Error())
	}

	if doc.Checksum != nil {
		err = doc.Checksum.Validate()
		if err != nil {
			return sdk.ErrUnknownRequest(err.Error())
		}
	}

	if doc.EncryptionData != nil {
		err = doc.EncryptionData.Validate()
		if err != nil {
			return sdk.ErrUnknownRequest(err.Error())
		}
	}

	if doc.EncryptionData != nil {

		for _, recipient := range doc.Recipients {
			found := false

			// Check that each address inside the EncryptionData object is contained inside the list of addresses
			for _, encAdd := range doc.EncryptionData.Keys {

				// Check that each recipient has an encrypted data associated to it
				if recipient.Equals(encAdd.Recipient) {
					found = true
				}

				if !doc.Recipients.Contains(encAdd.Recipient) {
					errMsg := fmt.Sprintf(
						"%s is a recipient inside encryption data but not inside the message",
						encAdd.Recipient.String(),
					)
					return sdk.ErrInvalidAddress(errMsg)
				}
			}

			if !found {
				// The recipient is not found inside the list of encrypted data recipients
				errMsg := fmt.Sprintf("%s is a recipient but has no encryption data specified", recipient.String())
				return sdk.ErrInvalidAddress(errMsg)
			}
		}

	}

	return nil
}

type Documents []Document

func (documents Documents) AppendIfMissing(i Document) Documents {
	for _, ele := range documents {
		if ele.Equals(i) {
			return documents
		}
	}
	return append(documents, i)
}

func (documents Documents) IsEmpty() bool {
	return len(documents) == 0
}
