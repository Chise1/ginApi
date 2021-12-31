package openapi

import "github.com/ginApi/openapi/models"

var OPENAPI = models.OpenAPI{}
var DefaultRes = models.Response{
	Description: "Successful Response",
	Content: map[string]models.MediaType{
		"application/json": {
			Schema: models.Schema{},
		},
	},
}

func init() {
	OPENAPI.Openapi = "3.0.2"
	OPENAPI.Info = models.Info{
		Title:   "UniBGP",
		Version: "0.1.0",
	}
	OPENAPI.Paths = map[string]models.PathItem{}
	OPENAPI.Components.Schemas = map[string]models.Schema{}
}
