package converter

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Fetch struct {
}

const fetchPattern = "fetch('%s', { method: '%s' })"

func NewFetchConverter() *Fetch {
	return &Fetch{}
}

func (c *Fetch) ConvertEndpoint(swagger *openapi3.T, method, endpoint string) (string, error) {
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

func (c *Fetch) convert(method string, endpoint string, _ *openapi3.Operation) string {
	return fmt.Sprintf(fetchPattern, method, endpoint)
}
