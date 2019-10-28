package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DocumentReceipt contains the generic information about the proof that a shared document (identified by the DocumentUUID field)
// has been read by the recipient.
// It contains information about the sender of the receipt's proof, the recipient,
// the original shared Document transaction's hash, a uuid that identifies which document has been received and
// the proof that document been read by its recipient.
// To be valid, all the fields of document receipt can't be empty and in particular
// transaction hash need to be referred to an existing transaction on chain and the uuid associated with document
// has to be non-empty and unique.
type DocumentReceipt struct {
	UUID         string         `json:"uuid"`
	Sender       sdk.AccAddress `json:"sender"`
	Recipient    sdk.AccAddress `json:"recipient"`
	TxHash       string         `json:"tx_hash"`
	DocumentUUID string         `json:"document_uuid"`
	Proof        string         `json:"proof"`
}

func (receipt DocumentReceipt) Equals(rec DocumentReceipt) bool {
	if !receipt.Sender.Equals(rec.Sender) {
		return false
	}
	if !receipt.Recipient.Equals(rec.Recipient) {
		return false
	}
	if receipt.TxHash != rec.TxHash {
		return false
	}
	if receipt.DocumentUUID != rec.DocumentUUID {
		return false
	}
	if receipt.Proof != rec.Proof {
		return false
	}
	return true
}

type DocumentReceipts []DocumentReceipt

func (receipts DocumentReceipts) IsEmpty() bool {
	return len(receipts) == 0
}

func (receipts DocumentReceipts) AppendIfMissing(receipt DocumentReceipt) (DocumentReceipts, bool) {
	for _, ele := range receipts {
		if ele.Equals(receipt) {
			return receipts, false
		}
	}
	return append(receipts, receipt), true
}

func (receipts DocumentReceipts) AppendAllIfMissing(other DocumentReceipts) DocumentReceipts {
	result := receipts
	for _, receipt := range other {
		result, _ = result.AppendIfMissing(receipt)
	}
	return result
}

func (receipts DocumentReceipts) FindByDocumentID(docID string) DocumentReceipts {
	var foundReceipts DocumentReceipts
	for _, ele := range receipts {
		if ele.DocumentUUID == docID {
			foundReceipts = append(foundReceipts, ele)
		}
	}
	return foundReceipts
}
