package openapi

type ServerVariable struct {
	Enum        []string `json:"enum"`        //An enumeration of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
	Default     string   `json:"default"`     //REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. Note this behavior is different than the Schema Object's treatment of default values, because in those cases parameter values are optional. If the enum is defined, the value MUST exist in the enum's values.
	Description string   `json:"description"` //An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.
}
type Server struct {
	Url         string                    `json:"url"`         //REQUIRED. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenAPI document is being served. Variable substitutions will be made when a variable is named in {brackets}.
	Description string                    `json:"description"` //An optional string describing the host designated by the URL. CommonMark syntax MAY be used for rich text representation.
	Variables   map[string]ServerVariable `json:"variables"`   //A map between a variable name and its value. The value is used for substitution in the server's URL template.
}
