package openapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOpenAPI(t *testing.T) {
	Docs := OpenAPI{
		Openapi: "3.0.2",
		Info: Info{
			Title:   "FastAPI",
			Version: "0.1.0",
		},
		Paths: map[string]PathItem{
			"/user/{Name}": {
				Get: Operation{
					Summary:     "Read Root",
					OperationId: "read_root_user__Name__get",
					Parameters: []Parameter{
						{
							Required: true,
							Schema:   Schema{Title: "Name", Type: "string"},
							Name:     "Name",
							In:       "path",
						},
					},
					Responses: map[string]Response{
						"200": {
							Description: "Successful Response",
							Content: map[string]MediaType{
								"application/json": {
									Schema: Schema{},
								},
							},
						},
						"422": {
							Description: "Validation Error",
							Content: map[string]MediaType{
								"application/json": {
									Schema: Schema{
										Ref: "#/components/schemas/HTTPValidationError",
									},
								},
							},
						},
					},
				},
			},
		},
		Components: Components{
			Schemas: map[string]Schema{
				"HTTPValidationError": {
					Title: "HTTPValidationError",
					Type:  "object",
					Properties: map[string]Schema{
						"detail": {
							Title: "Detail",
							Type:  "array",
							Items: &Schema{
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
					Properties: map[string]Schema{
						"loc": {
							Title: "Location",
							Type:  "array",
							Items: &Schema{
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
			},
		},
	}
	r := StructToMap(Docs)
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}
