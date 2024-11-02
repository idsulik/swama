package converter

import (
	"fmt"
	"strings"

	"github.com/idsulik/swama/v2/internal/model"
)

type Fetch struct {
}

const fetchPattern = "fetch('%s', { method: '%s', headers: %s, body: %s })"

func NewFetchConverter() *Fetch {
	return &Fetch{}
}

func (c *Fetch) ConvertEndpoint(method string, endpoint string, operation *model.Operation) string {
	headers := make(map[string]string)

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

			headers[param.Name] = value
		} else if param.In == model.ParameterInCookie {
			value := askForValue(param)

			if value == "" {
				continue
			}

			headers["Cookie"] = fmt.Sprintf("%s=%s", param.Name, value)
		}
	}

	var body string
	if operation.RequestBody != nil {
		contentType := askForContentType(operation.RequestBody.Value.Content)
		if contentType != "" {
			headers["Content-Type"] = contentType
		}

		body = "''"
	}

	headersBuilder := strings.Builder{}
	headersBuilder.WriteString("{")
	for k, v := range headers {
		headersBuilder.WriteString(fmt.Sprintf("'%s': '%s',", k, v))
	}
	headersBuilder.WriteString("}")

	return fmt.Sprintf(fetchPattern, endpoint, strings.ToUpper(method), headersBuilder.String(), body)
}
