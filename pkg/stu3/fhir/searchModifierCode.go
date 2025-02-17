package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SearchModifierCode is documented here http://hl7.org/fhir/ValueSet/search-modifier-code
type SearchModifierCode int

const (
	SearchModifierCodeMissing SearchModifierCode = iota
	SearchModifierCodeExact
	SearchModifierCodeContains
	SearchModifierCodeNot
	SearchModifierCodeText
	SearchModifierCodeIn
	SearchModifierCodeNotIn
	SearchModifierCodeBelow
	SearchModifierCodeAbove
	SearchModifierCodeType
)

func (code SearchModifierCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SearchModifierCode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "missing":
		*code = SearchModifierCodeMissing
	case "exact":
		*code = SearchModifierCodeExact
	case "contains":
		*code = SearchModifierCodeContains
	case "not":
		*code = SearchModifierCodeNot
	case "text":
		*code = SearchModifierCodeText
	case "in":
		*code = SearchModifierCodeIn
	case "not-in":
		*code = SearchModifierCodeNotIn
	case "below":
		*code = SearchModifierCodeBelow
	case "above":
		*code = SearchModifierCodeAbove
	case "type":
		*code = SearchModifierCodeType
	default:
		return fmt.Errorf("unknown SearchModifierCode code `%s`", s)
	}
	return nil
}
func (code SearchModifierCode) String() string {
	return code.Code()
}
func (code SearchModifierCode) Code() string {
	switch code {
	case SearchModifierCodeMissing:
		return "missing"
	case SearchModifierCodeExact:
		return "exact"
	case SearchModifierCodeContains:
		return "contains"
	case SearchModifierCodeNot:
		return "not"
	case SearchModifierCodeText:
		return "text"
	case SearchModifierCodeIn:
		return "in"
	case SearchModifierCodeNotIn:
		return "not-in"
	case SearchModifierCodeBelow:
		return "below"
	case SearchModifierCodeAbove:
		return "above"
	case SearchModifierCodeType:
		return "type"
	}
	return "<unknown>"
}
func (code SearchModifierCode) Display() string {
	switch code {
	case SearchModifierCodeMissing:
		return "Missing"
	case SearchModifierCodeExact:
		return "Exact"
	case SearchModifierCodeContains:
		return "Contains"
	case SearchModifierCodeNot:
		return "Not"
	case SearchModifierCodeText:
		return "Text"
	case SearchModifierCodeIn:
		return "In"
	case SearchModifierCodeNotIn:
		return "Not In"
	case SearchModifierCodeBelow:
		return "Below"
	case SearchModifierCodeAbove:
		return "Above"
	case SearchModifierCodeType:
		return "Type"
	}
	return "<unknown>"
}
func (code SearchModifierCode) Definition() string {
	switch code {
	case SearchModifierCodeMissing:
		return "The search parameter returns resources that have a value or not."
	case SearchModifierCodeExact:
		return "The search parameter returns resources that have a value that exactly matches the supplied parameter (the whole string, including casing and accents)."
	case SearchModifierCodeContains:
		return "The search parameter returns resources that include the supplied parameter value anywhere within the field being searched."
	case SearchModifierCodeNot:
		return "The search parameter returns resources that do not contain a match."
	case SearchModifierCodeText:
		return "The search parameter is processed as a string that searches text associated with the code/value - either CodeableConcept.text, Coding.display, or Identifier.type.text."
	case SearchModifierCodeIn:
		return "The search parameter is a URI (relative or absolute) that identifies a value set, and the search parameter tests whether the coding is in the specified value set."
	case SearchModifierCodeNotIn:
		return "The search parameter is a URI (relative or absolute) that identifies a value set, and the search parameter tests whether the coding is not in the specified value set."
	case SearchModifierCodeBelow:
		return "The search parameter tests whether the value in a resource is subsumed by the specified value (is-a, or hierarchical relationships)."
	case SearchModifierCodeAbove:
		return "The search parameter tests whether the value in a resource subsumes the specified value (is-a, or hierarchical relationships)."
	case SearchModifierCodeType:
		return "The search parameter only applies to the Resource Type specified as a modifier (e.g. the modifier is not actually :type, but :Patient etc.)."
	}
	return "<unknown>"
}
