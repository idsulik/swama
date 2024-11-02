package converter

import (
	"fmt"
	"strings"

	"github.com/idsulik/swama/v2/internal/model"
)

type Curl struct {
}

const curlPattern = "curl -X %s %s"

func NewCurlConverter() *Curl {
	return &Curl{}
}

func (c *Curl) ConvertEndpoint(method string, endpoint string, operation *model.Operation) string {
	var headers string

	for _, param := range operation.Parameters {
		if param.In == model.ParameterInPath {
			value := askForValue(param)

			if value == "" {
				continue
			}

			endpoint = strings.Replace(endpoint, fmt.Sprintf("{%s}", param.Name), value, 1)
		} else if param.In == model.ParameterInQuery {
			value := askForValue(param)

			if value == "" {
				continue
			}

			if strings.Contains(endpoint, "?") {
				endpoint = fmt.Sprintf("%s&%s=%s", endpoint, param.Name, value)
			} else {
				endpoint = fmt.Sprintf("%s?%s=%s", endpoint, param.Name, value)
			}
		} else if param.In == model.ParameterInHeader {
			value := askForValue(param)

			if value == "" {
				continue
			}

			headers += fmt.Sprintf(" -H '%s: %s'", param.Name, value)
		} else if param.In == model.ParameterInCookie {
			value := askForValue(param)

			if value == "" {
				continue
			}

			headers += fmt.Sprintf(" -H 'Cookie: %s=%s'", param.Name, value)
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
