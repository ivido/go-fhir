package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// RequestIntent is documented here http://hl7.org/fhir/ValueSet/request-intent
type RequestIntent int

const (
	RequestIntentProposal RequestIntent = iota
	RequestIntentPlan
	RequestIntentOrder
	RequestIntentOriginalOrder
	RequestIntentReflexOrder
	RequestIntentFillerOrder
	RequestIntentInstanceOrder
	RequestIntentOption
)

func (code RequestIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RequestIntent) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposal":
		*code = RequestIntentProposal
	case "plan":
		*code = RequestIntentPlan
	case "order":
		*code = RequestIntentOrder
	case "original-order":
		*code = RequestIntentOriginalOrder
	case "reflex-order":
		*code = RequestIntentReflexOrder
	case "filler-order":
		*code = RequestIntentFillerOrder
	case "instance-order":
		*code = RequestIntentInstanceOrder
	case "option":
		*code = RequestIntentOption
	default:
		return fmt.Errorf("unknown RequestIntent code `%s`", s)
	}
	return nil
}
func (code RequestIntent) String() string {
	return code.Code()
}
func (code RequestIntent) Code() string {
	switch code {
	case RequestIntentProposal:
		return "proposal"
	case RequestIntentPlan:
		return "plan"
	case RequestIntentOrder:
		return "order"
	case RequestIntentOriginalOrder:
		return "original-order"
	case RequestIntentReflexOrder:
		return "reflex-order"
	case RequestIntentFillerOrder:
		return "filler-order"
	case RequestIntentInstanceOrder:
		return "instance-order"
	case RequestIntentOption:
		return "option"
	}
	return "<unknown>"
}
func (code RequestIntent) Display() string {
	switch code {
	case RequestIntentProposal:
		return "Proposal"
	case RequestIntentPlan:
		return "Plan"
	case RequestIntentOrder:
		return "Order"
	case RequestIntentOriginalOrder:
		return "Original Order"
	case RequestIntentReflexOrder:
		return "Reflex Order"
	case RequestIntentFillerOrder:
		return "Filler Order"
	case RequestIntentInstanceOrder:
		return "Instance Order"
	case RequestIntentOption:
		return "Option"
	}
	return "<unknown>"
}
func (code RequestIntent) Definition() string {
	switch code {
	case RequestIntentProposal:
		return "The request is a suggestion made by someone/something that doesn't have an intention to ensure it occurs and without providing an authorization to act"
	case RequestIntentPlan:
		return "The request represents an intension to ensure something occurs without providing an authorization for others to act"
	case RequestIntentOrder:
		return "The request represents a request/demand and authorization for action"
	case RequestIntentOriginalOrder:
		return "The request represents an original authorization for action"
	case RequestIntentReflexOrder:
		return "The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization"
	case RequestIntentFillerOrder:
		return "The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order"
	case RequestIntentInstanceOrder:
		return "An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug."
	case RequestIntentOption:
		return "The request represents a component or option for a RequestGroup that establishes timing, conditionality and/or other constraints among a set of requests.\n\nRefer to [[[RequestGroup]]] for additional information on how this status is used"
	}
	return "<unknown>"
}
