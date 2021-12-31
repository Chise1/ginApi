package models

type Discriminator struct {
	PropertyName string            `json:"propertyName" gA:"need"` //	REQUIRED. The name of the property in the payload that will hold the discriminator value.
	Mapping      map[string]string `json:"mapping"`                //An object to hold mappings between payload values and schema names or references.
}
type XML struct {
	Name      string `json:"name"`      //	Replaces the name of the element/attribute used for the described schema property. When defined within items, it will affect the name of the individual XML elements within the list. When defined alongside type being array (outside the items), it will affect the wrapping element and only if wrapped is true. If wrapped is false, it will be ignored.
	Namespace string `json:"namespace"` // The URI of the namespace definition. This MUST be in the form of an absolute URI.
	Prefix    string `json:"prefix"`    //	The prefix to be used for the name.
	Attribute bool   `json:"attribute"` //	Declares whether the property definition translates to an attribute instead of an element. Default value is false.
	Wrapped   bool   `json:"wrapped"`   //	MAY be used only for an array definition. Signifies whether the array is wrapped (for example, <books><book/><book/></books>) or unwrapped (<book/><book/>). Default value is false. The definition takes effect only when defined alongside type being array (outside the items).
}

type Schema struct {
	Ref                  string                `json:"$ref"`
	Title                string                `json:"title"`
	MultipleOf           float64               `json:"multipleOf"`
	Maximum              float64               `json:"maximum"`
	ExclusiveMaximum     float64               `json:"exclusiveMaximum"`
	Minimum              float64               `json:"minimum"`
	ExclusiveMinimum     float64               `json:"exclusiveMinimum"`
	MaxLength            uint32                `json:"maxLength"`
	MinLength            uint32                `json:"minLength"`
	Pattern              string                `json:"pattern"`
	MaxItems             uint32                `json:"maxItems"`
	MinItems             uint32                `json:"minItems"`
	UniqueItems          bool                  `json:"uniqueItems"`
	MaxProperties        uint32                `json:"maxProperties"`
	MinProperties        uint32                `json:"minProperties"`
	Required             []string              `json:"required"`
	Enum                 []interface{}         `json:"enum"`
	Type                 string                `json:"type"`
	AllOf                []Schema              `json:"allOf"`
	OneOf                []Schema              `json:"oneOf"`
	AnyOf                []Schema              `json:"anyOf"`
	Not                  *Schema               `json:"not"`
	Items                *Schema               `json:"items"`
	Properties           map[string]Schema     `json:"properties"`
	AdditionalProperties interface{}           `json:"additionalProperties"`
	Description          string                `json:"description"`
	Format               string                `json:"format"`
	Default              interface{}           `json:"default"`
	Nullable             bool                  `json:"nullable"`
	ReadOnly             bool                  `json:"readOnly"`
	WriteOnly            bool                  `json:"writeOnly"`
	Discriminator        Discriminator         `json:"discriminator"` //    Adds support for polymorphism.The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description.See Composition and Inheritance for more details.
	Xml                  XML                   `json:"xml"`           //    This MAY be used only on properties schemas.It has no effect on root schemas.Adds additional metadata to describe the XML representation of this property.
	ExternalDocs         ExternalDocumentation `json:"externalDocs"`  //   Additional external documentation for this schema.
	Example              interface{}           `json:"example"`
	Deprecated           bool                  `json:"deprecated"`
	*Reference
}

func (n *Schema) SchemaSetAdditionalProperties(schema Schema) {
	n.AdditionalProperties = schema
}
func (n *Schema) ReferenceSetAdditionalProperties(schema Reference) {
	n.AdditionalProperties = schema
}
func (n *Schema) BoolSetAdditionalProperties(schema bool) {
	n.AdditionalProperties = schema
}
