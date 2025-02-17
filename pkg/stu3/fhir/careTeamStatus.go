package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CareTeamStatus is documented here http://hl7.org/fhir/ValueSet/care-team-status
type CareTeamStatus int

const (
	CareTeamStatusProposed CareTeamStatus = iota
	CareTeamStatusActive
	CareTeamStatusSuspended
	CareTeamStatusInactive
	CareTeamStatusEnteredInError
)

func (code CareTeamStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CareTeamStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposed":
		*code = CareTeamStatusProposed
	case "active":
		*code = CareTeamStatusActive
	case "suspended":
		*code = CareTeamStatusSuspended
	case "inactive":
		*code = CareTeamStatusInactive
	case "entered-in-error":
		*code = CareTeamStatusEnteredInError
	default:
		return fmt.Errorf("unknown CareTeamStatus code `%s`", s)
	}
	return nil
}
func (code CareTeamStatus) String() string {
	return code.Code()
}
func (code CareTeamStatus) Code() string {
	switch code {
	case CareTeamStatusProposed:
		return "proposed"
	case CareTeamStatusActive:
		return "active"
	case CareTeamStatusSuspended:
		return "suspended"
	case CareTeamStatusInactive:
		return "inactive"
	case CareTeamStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code CareTeamStatus) Display() string {
	switch code {
	case CareTeamStatusProposed:
		return "Proposed"
	case CareTeamStatusActive:
		return "Active"
	case CareTeamStatusSuspended:
		return "Suspended"
	case CareTeamStatusInactive:
		return "Inactive"
	case CareTeamStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code CareTeamStatus) Definition() string {
	switch code {
	case CareTeamStatusProposed:
		return "The care team has been drafted and proposed, but not yet participating in the coordination and delivery of care."
	case CareTeamStatusActive:
		return "The care team is currently participating in the coordination and delivery of care."
	case CareTeamStatusSuspended:
		return "The care team is temporarily on hold or suspended and not participating in the coordination and delivery of care."
	case CareTeamStatusInactive:
		return "The care team was, but is no longer, participating in the coordination and delivery of care."
	case CareTeamStatusEnteredInError:
		return "The care team should have never existed."
	}
	return "<unknown>"
}
