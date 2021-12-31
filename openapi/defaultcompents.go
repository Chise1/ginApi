package openapi

import "github.com/ginApi/openapi/models"

// DefaultValidationCompents 默认的报错结构
func DefaultValidationCompents() map[string]models.Schema {
	return map[string]models.Schema{
		"HTTPValidationError": {
			Title: "HTTPValidationError",
			Type:  "object",
			Properties: map[string]models.Schema{
				"detail": {
					Title: "Detail",
					Type:  "array",
					Items: &models.Schema{
						Ref: "#/components/schemas/ValidationError",
					},
				},
			},
		},
		"ValidationError": {
			Title: "ValidationError",
			Required: []string{
				"loc", "msg", "type",
			},
			Type: "object",
			Properties: map[string]models.Schema{
				"loc": {
					Title: "Location",
					Type:  "array",
					Items: &models.Schema{
						Type: "string",
					},
				},
				"msg": {
					Title: "Message",
					Type:  "string",
				},
				"type": {
					Title: "Error Type",
					Type:  "string",
				},
			},
		},
	}
}

//默认的报错返回
func DefaultResponses() map[string]models.Response {
	return map[string]models.Response{
		"200": {
			Description: "Successful Response",
			Content: map[string]models.MediaType{
				"application/json": {
					Schema: models.Schema{},
				},
			},
		},
		"422": {
			Description: "Validation Error",
			Content: map[string]models.MediaType{
				"application/json": {
					Schema: models.Schema{
						Ref: "#/components/schemas/HTTPValidationError",
					},
				},
			},
		},
	}
}
