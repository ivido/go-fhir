package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// QuestionnaireResponseStatus is documented here http://hl7.org/fhir/ValueSet/questionnaire-answers-status
type QuestionnaireResponseStatus int

const (
	QuestionnaireResponseStatusInProgress QuestionnaireResponseStatus = iota
	QuestionnaireResponseStatusCompleted
	QuestionnaireResponseStatusAmended
	QuestionnaireResponseStatusEnteredInError
	QuestionnaireResponseStatusStopped
)

func (code QuestionnaireResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuestionnaireResponseStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "in-progress":
		*code = QuestionnaireResponseStatusInProgress
	case "completed":
		*code = QuestionnaireResponseStatusCompleted
	case "amended":
		*code = QuestionnaireResponseStatusAmended
	case "entered-in-error":
		*code = QuestionnaireResponseStatusEnteredInError
	case "stopped":
		*code = QuestionnaireResponseStatusStopped
	default:
		return fmt.Errorf("unknown QuestionnaireResponseStatus code `%s`", s)
	}
	return nil
}
func (code QuestionnaireResponseStatus) String() string {
	return code.Code()
}
func (code QuestionnaireResponseStatus) Code() string {
	switch code {
	case QuestionnaireResponseStatusInProgress:
		return "in-progress"
	case QuestionnaireResponseStatusCompleted:
		return "completed"
	case QuestionnaireResponseStatusAmended:
		return "amended"
	case QuestionnaireResponseStatusEnteredInError:
		return "entered-in-error"
	case QuestionnaireResponseStatusStopped:
		return "stopped"
	}
	return "<unknown>"
}
func (code QuestionnaireResponseStatus) Display() string {
	switch code {
	case QuestionnaireResponseStatusInProgress:
		return "In Progress"
	case QuestionnaireResponseStatusCompleted:
		return "Completed"
	case QuestionnaireResponseStatusAmended:
		return "Amended"
	case QuestionnaireResponseStatusEnteredInError:
		return "Entered in Error"
	case QuestionnaireResponseStatusStopped:
		return "Stopped"
	}
	return "<unknown>"
}
func (code QuestionnaireResponseStatus) Definition() string {
	switch code {
	case QuestionnaireResponseStatusInProgress:
		return "This QuestionnaireResponse has been partially filled out with answers, but changes or additions are still expected to be made to it."
	case QuestionnaireResponseStatusCompleted:
		return "This QuestionnaireResponse has been filled out with answers, and the current content is regarded as definitive."
	case QuestionnaireResponseStatusAmended:
		return "This QuestionnaireResponse has been filled out with answers, then marked as complete, yet changes or additions have been made to it afterwards."
	case QuestionnaireResponseStatusEnteredInError:
		return "This QuestionnaireResponse was entered in error and voided."
	case QuestionnaireResponseStatusStopped:
		return "This QuestionnaireResponse has been partially filled out with answers, but has been abandoned. It is unknown whether changes or additions are expected to be made to it."
	}
	return "<unknown>"
}
