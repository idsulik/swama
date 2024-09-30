package model

import "github.com/getkin/kin-openapi/openapi3"

type Operation struct {
	Method string
	Path   string
	*openapi3.Operation
}

func NewOperation(method, path string, operation *openapi3.Operation) *Operation {
	return &Operation{
		Method:    method,
		Path:      path,
		Operation: operation,
	}
}
