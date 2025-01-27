package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AllergyIntoleranceType is documented here http://hl7.org/fhir/ValueSet/allergy-intolerance-type
type AllergyIntoleranceType int

const (
	AllergyIntoleranceTypeAllergy AllergyIntoleranceType = iota
	AllergyIntoleranceTypeIntolerance
)

func (code AllergyIntoleranceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "allergy":
		*code = AllergyIntoleranceTypeAllergy
	case "intolerance":
		*code = AllergyIntoleranceTypeIntolerance
	default:
		return fmt.Errorf("unknown AllergyIntoleranceType code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceType) String() string {
	return code.Code()
}
func (code AllergyIntoleranceType) Code() string {
	switch code {
	case AllergyIntoleranceTypeAllergy:
		return "allergy"
	case AllergyIntoleranceTypeIntolerance:
		return "intolerance"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceType) Display() string {
	switch code {
	case AllergyIntoleranceTypeAllergy:
		return "Allergy"
	case AllergyIntoleranceTypeIntolerance:
		return "Intolerance"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceType) Definition() string {
	switch code {
	case AllergyIntoleranceTypeAllergy:
		return "A propensity for hypersensitivity reaction(s) to a substance.  These reactions are most typically type I hypersensitivity, plus other \"allergy-like\" reactions, including pseudoallergy."
	case AllergyIntoleranceTypeIntolerance:
		return "A propensity for adverse reactions to a substance that is not judged to be allergic or \"allergy-like\".  These reactions are typically (but not necessarily) non-immune.  They are to some degree idiosyncratic and/or individually specific (i.e. are not a reaction that is expected to occur with most or all patients given similar circumstances)."
	}
	return "<unknown>"
}
