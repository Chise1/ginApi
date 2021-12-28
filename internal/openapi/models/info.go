package models

type Contact struct {
	Name  string `json:"name"`  //The identifying name of the contact person/organization.
	Url   string `json:"url"`   //The URL pointing to the contact information. This MUST be in the form of a URL.
	Email string `json:"email"` //The email address of the contact person/organization. This MUST be in the form of an email address.
}

type License struct {
	Name string `json:"name"` //REQUIRED. The license name used for the API.
	Url  string `json:"url"`  //A URL to the license used for the API. This MUST be in the form of a URL. The url field is mutually exclusive of the identifier field.
}

type Info struct {
	Title          string  `json:"title"`          //REQUIRED. The title of the API.
	Summary        string  `json:"summary"`        //A short summary of the API.
	Description    string  `json:"description"`    //A description of the API. CommonMark syntax MAY be used for rich text representation.
	TermsOfService string  `json:"termsOfService"` //A URL to the Terms of Service for the API. This MUST be in the form of a URL.
	Contact        Contact `json:"contact"`        //The contact information for the exposed API.
	License        License `json:"license"`        //The license information for the exposed API.
	Version        string  `json:"version"`        //REQUIRED. The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).
}
