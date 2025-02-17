package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// VisionBase is documented here http://hl7.org/fhir/ValueSet/vision-base-codes
type VisionBase int

const (
	VisionBaseUp VisionBase = iota
	VisionBaseDown
	VisionBaseIn
	VisionBaseOut
)

func (code VisionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *VisionBase) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "up":
		*code = VisionBaseUp
	case "down":
		*code = VisionBaseDown
	case "in":
		*code = VisionBaseIn
	case "out":
		*code = VisionBaseOut
	default:
		return fmt.Errorf("unknown VisionBase code `%s`", s)
	}
	return nil
}
func (code VisionBase) String() string {
	return code.Code()
}
func (code VisionBase) Code() string {
	switch code {
	case VisionBaseUp:
		return "up"
	case VisionBaseDown:
		return "down"
	case VisionBaseIn:
		return "in"
	case VisionBaseOut:
		return "out"
	}
	return "<unknown>"
}
func (code VisionBase) Display() string {
	switch code {
	case VisionBaseUp:
		return "Up"
	case VisionBaseDown:
		return "Down"
	case VisionBaseIn:
		return "In"
	case VisionBaseOut:
		return "Out"
	}
	return "<unknown>"
}
func (code VisionBase) Definition() string {
	switch code {
	case VisionBaseUp:
		return "top"
	case VisionBaseDown:
		return "bottom"
	case VisionBaseIn:
		return "inner edge"
	case VisionBaseOut:
		return "outer edge"
	}
	return "<unknown>"
}
