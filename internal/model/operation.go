package model

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Operation struct {
	Method     string
	Path       string
	Parameters []*Parameter
	*openapi3.Operation
}

func NewOperation(method, path string, operation *openapi3.Operation) *Operation {
	return &Operation{
		Method:     method,
		Path:       path,
		Parameters: createParameters(operation.Parameters),
		Operation:  operation,
	}
}

func createParameters(parameters openapi3.Parameters) []*Parameter {
	params := make([]*Parameter, 0, len(parameters))
	for _, p := range parameters {
		propertyType := ""
		if p.Value.Schema != nil {
			propertyType = strings.Join(p.Value.Schema.Value.Type.Slice(), ", ")
		}

		params = append(
			params,
			NewParameter(p.Value.In, p.Value.Name, propertyType, p.Value.Required, p.Value.Description),
		)
	}
	return params
}
