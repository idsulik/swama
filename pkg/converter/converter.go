package converter

import (
	"errors"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

var (
	ErrEndpointNotFound    = fmt.Errorf("endpoint not found")
	ErrEndpointWrongMethod = fmt.Errorf("endpoint found but wrong method")
)

type Converter interface {
	ConvertEndpoint(swagger *openapi3.T, method, endpoint string) (string, error)
}

const (
	CurlType  = "curl"
	FetchType = "fetch"
)

var (
	ErrInvalidConvertType = errors.New(fmt.Sprintf("invalid convert type. Must be %q or %q", CurlType, FetchType))
)

func NewConverter(convertType string) (Converter, error) {
	switch convertType {
	case CurlType:
		return NewCurlConverter(), nil
	case FetchType:
		return NewFetchConverter(), nil
	default:
		return nil, ErrInvalidConvertType
	}
}
