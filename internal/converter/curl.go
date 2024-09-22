package converter

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Curl struct {
}

const curlPattern = "curl -X %s %s"

func NewCurlConverter() *Curl {
	return &Curl{}
}

func (c *Curl) ConvertEndpoint(method string, endpoint string, operation *openapi3.Operation) string {
	var headers string

	for _, param := range operation.Parameters {
		if param.Value.In == openapi3.ParameterInPath {
			value := askForValue(param.Value)

			if value == "" {
				continue
			}

			endpoint = strings.Replace(endpoint, fmt.Sprintf("{%s}", param.Value.Name), value, 1)
		} else if param.Value.In == openapi3.ParameterInQuery {
			value := askForValue(param.Value)

			if value == "" {
				continue
			}

			if strings.Contains(endpoint, "?") {
				endpoint = fmt.Sprintf("%s&%s=%s", endpoint, param.Value.Name, value)
			} else {
				endpoint = fmt.Sprintf("%s?%s=%s", endpoint, param.Value.Name, value)
			}
		} else if param.Value.In == openapi3.ParameterInHeader {
			value := askForValue(param.Value)

			if value == "" {
				continue
			}

			headers += fmt.Sprintf(" -H '%s: %s'", param.Value.Name, value)
		} else if param.Value.In == openapi3.ParameterInCookie {
			value := askForValue(param.Value)

			if value == "" {
				continue
			}

			headers += fmt.Sprintf(" -H 'Cookie: %s=%s'", param.Value.Name, value)
		}
	}

	var body string
	if operation.RequestBody != nil {
		contentType := askForContentType(operation.RequestBody.Value.Content)
		if contentType != "" {
			headers += fmt.Sprintf(" -H 'Content-Type: %s'", contentType)
		}

		body = " -d ''"
	}

	return fmt.Sprintf(curlPattern, strings.ToUpper(method), endpoint+headers+body)
}
