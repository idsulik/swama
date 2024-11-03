package mockserver

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

const (
	responseCodeHeaderName       = "X-Mock-Response-Code"
	responseCodeQueryParamName   = "x-response-code"
	responseTypeQueryParamName   = "x-response-type"
	availableResponsesHeaderName = "X-Available-Responses"
)

type RunOptions struct {
	Host                string
	Port                int
	Delay               int
	DefaultResponseCode string
	DefaultResponseType string
}

// XMLNode represents a generic XML node
type XMLNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:"attr,omitempty"`
	Value    string     `xml:",chardata"`
	Children []*XMLNode `xml:",any"`
}

type MockServer struct {
	doc *openapi3.T
}

func NewMockServer(doc *openapi3.T) *MockServer {
	return &MockServer{
		doc: doc,
	}
}

func (m *MockServer) Run(options RunOptions) error {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.Use(m.availableResponsesMiddleware())

	for _, path := range m.doc.Paths.InMatchingOrder() {
		for method, operation := range m.doc.Paths.Find(path).Operations() {
			app.Handle(method, convertPathToGinFormat(path), m.registerHandler(operation, options))
		}
	}

	// Index route to list all registered routes
	app.GET(
		"/", func(c *gin.Context) {
			var routes []gin.H
			for _, route := range app.Routes() {
				if route.Path == "/" {
					continue
				}

				path := m.doc.Paths.Find(convertGinPathToOpenAPI(route.Path))
				responseCodeToContentType := make(map[string]string)
				if op := path.GetOperation(route.Method); op != nil {
					for code, resp := range op.Responses.Map() {
						responseCodeToContentType[code] = "application/json"
						for contentType := range resp.Value.Content {
							responseCodeToContentType[code] = contentType
						}
					}
				}

				routes = append(
					routes, gin.H{
						"method":             route.Method,
						"path":               route.Path,
						"availableResponses": responseCodeToContentType,
					},
				)
			}

			c.JSON(
				http.StatusOK, gin.H{
					"routes": routes,
					"usage": gin.H{
						"responseCode": gin.H{
							"queryParam":     fmt.Sprintf("?%s=<status_code>", responseCodeQueryParamName),
							"header":         fmt.Sprintf("%s: <status_code>", responseCodeHeaderName),
							"availableCodes": fmt.Sprintf("%s header in response", availableResponsesHeaderName),
						},
						"responseType": gin.H{
							"queryParam": fmt.Sprintf("?%s=<response_type>", responseTypeQueryParamName),
						},
					},
				},
			)
		},
	)

	fmt.Printf("Mock server listening on http://%s:%d\n", options.Host, options.Port)
	return app.Run(fmt.Sprintf("%s:%d", options.Host, options.Port))
}

func (m *MockServer) availableResponsesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// After handler execution, add header with available response codes
		path := m.doc.Paths.Find(convertGinPathToOpenAPI(c.FullPath()))
		if path != nil {
			if op := path.GetOperation(c.Request.Method); op != nil {
				var codes []string
				for code := range op.Responses.Map() {
					codes = append(codes, code)
				}
				if len(codes) > 0 {
					c.Header(availableResponsesHeaderName, strings.Join(codes, ","))
				}
			}
		}
	}
}

func (m *MockServer) registerHandler(op *openapi3.Operation, options RunOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		// If Delay is set in options, apply it to simulate latency
		if options.Delay > 0 {
			time.Sleep(time.Duration(options.Delay) * time.Millisecond)
		}

		// Get desired response code from query param or header
		desiredCode := c.Query(responseCodeQueryParamName)
		if desiredCode == "" {
			desiredCode = c.GetHeader(responseCodeHeaderName)
		}

		// If no code specified, use default
		if desiredCode == "" {
			desiredCode = options.DefaultResponseCode
		}

		// Get desired response type from query param
		desiredType := c.Query(responseTypeQueryParamName)
		if desiredType == "" {
			desiredType = options.DefaultResponseType
		}

		response := m.findSpecificResponse(op, desiredCode)
		if response == nil {
			body := gin.H{
				"error":          fmt.Sprintf("No response defined for status code %s", desiredCode),
				"availableCodes": m.getAvailableResponseCodes(op),
			}
			if desiredType == "xml" {
				c.XML(http.StatusBadRequest, body)
				return
			} else {
				c.JSON(http.StatusBadRequest, body)
			}
			return
		}

		status := parseStatusCode(desiredCode)
		acceptHeader := c.GetHeader("Accept")
		acceptedTypes := parseAcceptHeader(acceptHeader)

		var contentType string
		var schema *openapi3.Schema
		for mediaType, content := range response.Content {
			for _, acceptedType := range acceptedTypes {
				if desiredType != "" {
					if strings.HasSuffix(mediaType, desiredType) {
						contentType = mediaType
						schema = content.Schema.Value
						break
					}

				} else if strings.HasPrefix(mediaType, acceptedType) {
					contentType = mediaType
					schema = content.Schema.Value
					break
				}
			}

			if schema != nil {
				break
			}
		}

		if schema == nil {
			contentType = "application/json"
			if jsonContent, ok := response.Content["application/json"]; ok {
				schema = jsonContent.Schema.Value
			}
		}

		mockData := generateMockData(schema)

		switch {
		case strings.Contains(contentType, "application/xml"):
			c.XML(status, mapToXML(mockData, schema, "root"))
		default:
			c.JSON(status, mockData)
		}
	}
}

func (m *MockServer) findSpecificResponse(op *openapi3.Operation, code string) *openapi3.Response {
	if responseRef, ok := op.Responses.Map()[code]; ok {
		return responseRef.Value
	}
	return nil
}

func (m *MockServer) getAvailableResponseCodes(op *openapi3.Operation) []string {
	var codes []string
	for code := range op.Responses.Map() {
		codes = append(codes, code)
	}
	return codes
}
