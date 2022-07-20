package fhir

import "encoding/json"

// StructureDefinition is documented here http://hl7.org/fhir/StructureDefinition/StructureDefinition
type StructureDefinition struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                           `bson:"url" json:"url"`
	Identifier        []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                          `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                           `bson:"name" json:"name"`
	Title             *string                          `bson:"title,omitempty" json:"title,omitempty"`
	Status            string                           `bson:"status" json:"status"`
	Experimental      *bool                            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                          `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                  `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                   `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                          `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Keyword           []Coding                         `bson:"keyword,omitempty" json:"keyword,omitempty"`
	FhirVersion       *string                          `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Mapping           []StructureDefinitionMapping     `bson:"mapping,omitempty" json:"mapping,omitempty"`
	Kind              string                           `bson:"kind" json:"kind"`
	Abstract          bool                             `bson:"abstract" json:"abstract"`
	ContextType       *string                          `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Context           []string                         `bson:"context,omitempty" json:"context,omitempty"`
	ContextInvariant  []string                         `bson:"contextInvariant,omitempty" json:"contextInvariant,omitempty"`
	Type              string                           `bson:"type" json:"type"`
	BaseDefinition    *string                          `bson:"baseDefinition,omitempty" json:"baseDefinition,omitempty"`
	Derivation        *string                          `bson:"derivation,omitempty" json:"derivation,omitempty"`
	Snapshot          *StructureDefinitionSnapshot     `bson:"snapshot,omitempty" json:"snapshot,omitempty"`
	Differential      *StructureDefinitionDifferential `bson:"differential,omitempty" json:"differential,omitempty"`
}
type StructureDefinitionMapping struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identity          string      `bson:"identity" json:"identity"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
type StructureDefinitionSnapshot struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `bson:"element" json:"element"`
}
type StructureDefinitionDifferential struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `bson:"element" json:"element"`
}
type OtherStructureDefinition StructureDefinition

// MarshalJSON marshals the given StructureDefinition as JSON into a byte slice
func (r StructureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureDefinition: OtherStructureDefinition(r),
		ResourceType:             "StructureDefinition",
	})
}

// UnmarshalStructureDefinition unmarshals a StructureDefinition.
func UnmarshalStructureDefinition(b []byte) (StructureDefinition, error) {
	var structureDefinition StructureDefinition
	if err := json.Unmarshal(b, &structureDefinition); err != nil {
		return structureDefinition, err
	}
	return structureDefinition, nil
}