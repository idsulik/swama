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

func (c *Curl) ConvertEndpoint(method string, endpoint string, _ *openapi3.Operation) string {
	return fmt.Sprintf(curlPattern, method, endpoint)
}
