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

func (c *Fetch) ConvertEndpoint(method string, endpoint string, _ *openapi3.Operation) string {
	return fmt.Sprintf(fetchPattern, method, endpoint)
}
