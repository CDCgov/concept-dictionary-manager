package main

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
	Oid              string `json:"oid"`
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

//Code System from ElasticSearch
type ESSystem struct {
	Oid                           string
	ID                            string
	Name                          string
	DefinitionText                string
	Status                        string
	StatusDate                    string
	Version                       string
	VersionDescription            string
	AcquiredDate                  string
	AssigningAuthorityVersionName string
	AssigningAuthorityReleaseDate string
	DistributionSourceVersionName string
	DistributionSourceReleaseDate string
	DistributionSourceID          string
	LastRevisionDate              string
	AssigningAuthorityID          string
	CodeSystemCode                string
	SourceURL                     string
	Hl70396Identifier             string
	LegacyFlag                    bool
}

//Code from ElasticSearch
type ESCode struct {
	ID                      string
	Name                    string
	CodeSystemOid           string
	ConceptCode             string
	SdoPreferredDesignation string
	DefinitionText          string
	SdoConceptRevisionDate  string
	Status                  string
	StatusDate              string
	SdoConceptStatus        string
	SdoConceptStatusDate    string
	IsRootFlag              bool
	IsConceptFlat           bool
}
