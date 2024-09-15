package converter

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Curl struct {
}

const curlPattern = "Curl -X %s %s"

func NewCurlConverter() *Curl {
	return &Curl{}
}

func (c *Curl) ConvertEndpoint(swagger *openapi3.T, method, endpoint string) (string, error) {
	path := swagger.Paths.Find(endpoint)

	if path == nil {
		return "", ErrEndpointNotFound
	}

	for m, operation := range path.Operations() {
		if m == method {
			return c.convert(method, endpoint, operation), nil
		}
	}

	return "", ErrEndpointWrongMethod
}

func (c *Curl) convert(method string, endpoint string, _ *openapi3.Operation) string {
	return fmt.Sprintf(curlPattern, method, endpoint)
}
