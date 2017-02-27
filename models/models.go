package models

//CodeSystem struct to represent supported code systems
type CodeSystem struct {
	URL              string `json:"url"`
	Version          string `json:"version"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	Publisher        string `json:"publisher"`
	Copyright        string `json:"copyright"`
	CaseSensitive    bool   `json:"caseSensitive"`
	HierarchyMeaning string `json:"hierarchyMeaning"`
	Compositional    bool   `json:"compositional"`
	Content          string `json:"content"`
	Count            int    `json:"count"`
}

//ValueSet - A set of codes drawn from one or more code systems. This is a proper subset of the [FHIR ValueSet](http://hl7.org/fhir/2016Sep/valueset.html) structure.
type ValueSet struct {
	URL         string       `json:"url"`
	Identifier  []Identifier `json:"identifier"`
	Version     string       `json:"version"`
	Name        string       `json:"name"`
	Status      string       `json:"status"`
	Description string       `json:"description"`
	Expansion   Expansion    `json:"expansion"`
}

//Identifier for ValueSet
type Identifier struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

//Expansion for ValueSet
type Expansion struct {
	Identifier string `json:"identifier"`
	Timestamp  string `json:"timestamp"`
	Contains   []Code `json:"contains"`
}

//Code for Expansion
type Code struct {
	System  string `json:"system"`
	Version string `json:"version"`
	Code    string `json:"code"`
	Display string `json:"display"`
}
