package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ListStatus is documented here http://hl7.org/fhir/ValueSet/list-status
type ListStatus int

const (
	ListStatusCurrent ListStatus = iota
	ListStatusRetired
	ListStatusEnteredInError
)

func (code ListStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ListStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "current":
		*code = ListStatusCurrent
	case "retired":
		*code = ListStatusRetired
	case "entered-in-error":
		*code = ListStatusEnteredInError
	default:
		return fmt.Errorf("unknown ListStatus code `%s`", s)
	}
	return nil
}
func (code ListStatus) String() string {
	return code.Code()
}
func (code ListStatus) Code() string {
	switch code {
	case ListStatusCurrent:
		return "current"
	case ListStatusRetired:
		return "retired"
	case ListStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ListStatus) Display() string {
	switch code {
	case ListStatusCurrent:
		return "Current"
	case ListStatusRetired:
		return "Retired"
	case ListStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code ListStatus) Definition() string {
	switch code {
	case ListStatusCurrent:
		return "The list is considered to be an active part of the patient's record."
	case ListStatusRetired:
		return "The list is \"old\" and should no longer be considered accurate or relevant."
	case ListStatusEnteredInError:
		return "The list was never accurate.  It is retained for medico-legal purposes only."
	}
	return "<unknown>"
}
