package openapi

type Reference struct {
	Ref string `json:"$ref"` //REQUIRED. The reference identifier. This MUST be in the form of a URI.
}

func (n Reference) GetRef() string {
	return n.Ref
}
