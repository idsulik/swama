package mockserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

type RunOptions struct {
	Host  string
	Port  int
	Delay int
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
	app := gin.Default()

	for _, path := range m.doc.Paths.InMatchingOrder() {
		for method, operation := range m.doc.Paths.Find(path).Operations() {
			app.Handle(method, convertPathToGinFormat(path), m.registerHandler(method, path, operation, options))
		}
	}

	// Index route to list all registered routes
	app.GET(
		"/", func(c *gin.Context) {
			var routes []gin.H
			for _, route := range app.Routes() {
				routes = append(
					routes, gin.H{
						"method": route.Method,
						"path":   route.Path,
					},
				)
			}
			c.JSON(http.StatusOK, routes)
		},
	)

	return app.Run(fmt.Sprintf("%s:%d", options.Host, options.Port))
}

func (m *MockServer) registerHandler(method, path string, op *openapi3.Operation, options RunOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		// If Delay is set in options, apply it to simulate latency
		if options.Delay > 0 {
			time.Sleep(time.Duration(options.Delay) * time.Millisecond)
		}

		// Choose a response status code from the operation's responses
		var status int
		var example interface{}
		response := m.findResponse(op)

		// Use the first example found in the response content
		if len(response.Content) > 0 {
			for _, mediaType := range response.Content {
				if mediaType.Example != nil {
					example = mediaType.Example
				} else if len(mediaType.Examples) > 0 {
					// If multiple examples exist, use the first one
					for _, ex := range mediaType.Examples {
						example = ex.Value.Value
						break
					}
				}
				break
			}
		}

		// If no example is found, fall back to a generic message
		if example == nil {
			params := c.Params
			paramsStr := strings.Builder{}
			for _, p := range params {
				paramsStr.WriteString(fmt.Sprintf("%s=%s ", p.Key, p.Value))
			}
			example = gin.H{
				"message": fmt.Sprintf("Mock response"),
				"method":  method,
				"path":    path,
				"params":  params,
			}
		}

		c.JSON(status, example)
	}
}

func (m *MockServer) findResponse(op *openapi3.Operation) *openapi3.Response {
	var response *openapi3.Response
	for statusCode, responseRef := range op.Responses.Map() {
		status := parseStatusCode(statusCode)

		if status >= 200 && status < 300 {
			return responseRef.Value
		}

		if response == nil {
			response = responseRef.Value
		}
	}

	return response
}

// Helper function to convert OpenAPI path format to Gin's format
func convertPathToGinFormat(path string) string {
	path = strings.ReplaceAll(path, "{", ":")
	path = strings.ReplaceAll(path, "}", "")

	return path
}

// Helper function to parse the status code string to an integer
func parseStatusCode(code string) int {
	status, err := strconv.Atoi(code)
	if err != nil {
		log.Printf("Invalid status code '%s', defaulting to 200", code)
		return 200
	}
	return status
}
