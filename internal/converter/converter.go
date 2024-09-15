package converter

import (
	"errors"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

const (
	CurlType  = "curl"
	FetchType = "fetch"
)

var (
	ErrInvalidConvertType = errors.New(fmt.Sprintf("invalid convert type. Must be %q or %q", CurlType, FetchType))
)

type Converter interface {
	ConvertEndpoint(method string, endpoint string, _ *openapi3.Operation) string
}

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
