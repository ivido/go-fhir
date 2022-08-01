package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// BundleType is documented here http://hl7.org/fhir/ValueSet/bundle-type
type BundleType int

const (
	BundleTypeDocument BundleType = iota
	BundleTypeMessage
	BundleTypeTransaction
	BundleTypeTransactionResponse
	BundleTypeBatch
	BundleTypeBatchResponse
	BundleTypeHistory
	BundleTypeSearchset
	BundleTypeCollection
)

func (code BundleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BundleType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "document":
		*code = BundleTypeDocument
	case "message":
		*code = BundleTypeMessage
	case "transaction":
		*code = BundleTypeTransaction
	case "transaction-response":
		*code = BundleTypeTransactionResponse
	case "batch":
		*code = BundleTypeBatch
	case "batch-response":
		*code = BundleTypeBatchResponse
	case "history":
		*code = BundleTypeHistory
	case "searchset":
		*code = BundleTypeSearchset
	case "collection":
		*code = BundleTypeCollection
	default:
		return fmt.Errorf("unknown BundleType code `%s`", s)
	}
	return nil
}
func (code BundleType) String() string {
	return code.Code()
}
func (code BundleType) Code() string {
	switch code {
	case BundleTypeDocument:
		return "document"
	case BundleTypeMessage:
		return "message"
	case BundleTypeTransaction:
		return "transaction"
	case BundleTypeTransactionResponse:
		return "transaction-response"
	case BundleTypeBatch:
		return "batch"
	case BundleTypeBatchResponse:
		return "batch-response"
	case BundleTypeHistory:
		return "history"
	case BundleTypeSearchset:
		return "searchset"
	case BundleTypeCollection:
		return "collection"
	}
	return "<unknown>"
}
func (code BundleType) Display() string {
	switch code {
	case BundleTypeDocument:
		return "Document"
	case BundleTypeMessage:
		return "Message"
	case BundleTypeTransaction:
		return "Transaction"
	case BundleTypeTransactionResponse:
		return "Transaction Response"
	case BundleTypeBatch:
		return "Batch"
	case BundleTypeBatchResponse:
		return "Batch Response"
	case BundleTypeHistory:
		return "History List"
	case BundleTypeSearchset:
		return "Search Results"
	case BundleTypeCollection:
		return "Collection"
	}
	return "<unknown>"
}
func (code BundleType) Definition() string {
	switch code {
	case BundleTypeDocument:
		return "The bundle is a document. The first resource is a Composition."
	case BundleTypeMessage:
		return "The bundle is a message. The first resource is a MessageHeader."
	case BundleTypeTransaction:
		return "The bundle is a transaction - intended to be processed by a server as an atomic commit."
	case BundleTypeTransactionResponse:
		return "The bundle is a transaction response. Because the response is a transaction response, the transaction has succeeded, and all responses are error free."
	case BundleTypeBatch:
		return "The bundle is a transaction - intended to be processed by a server as a group of actions."
	case BundleTypeBatchResponse:
		return "The bundle is a batch response. Note that as a batch, some responses may indicate failure and others success."
	case BundleTypeHistory:
		return "The bundle is a list of resources from a history interaction on a server."
	case BundleTypeSearchset:
		return "The bundle is a list of resources returned as a result of a search/query interaction, operation, or message."
	case BundleTypeCollection:
		return "The bundle is a set of resources collected into a single package for ease of distribution."
	}
	return "<unknown>"
}