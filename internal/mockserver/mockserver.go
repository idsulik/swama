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
			app.Handle(method, convertPathToGinFormat(path), m.registerHandler(operation, options))
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

func (m *MockServer) registerHandler(op *openapi3.Operation, options RunOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		// If Delay is set in options, apply it to simulate latency
		if options.Delay > 0 {
			time.Sleep(time.Duration(options.Delay) * time.Millisecond)
		}

		status, response := m.findResponse(op)
		if response == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No response defined"})
			return
		}

		// Get accepted content types from Accept header
		acceptHeader := c.GetHeader("Accept")
		acceptedTypes := parseAcceptHeader(acceptHeader)

		// Find matching content type and schema
		var contentType string
		var schema *openapi3.Schema
		for mediaType, content := range response.Content {
			for _, acceptedType := range acceptedTypes {
				if strings.HasPrefix(mediaType, acceptedType) {
					contentType = mediaType
					schema = content.Schema.Value
					break
				}
			}
			if schema != nil {
				break
			}
		}

		// If no matching content type found, default to JSON
		if schema == nil {
			contentType = "application/json"
			if jsonContent, ok := response.Content["application/json"]; ok {
				schema = jsonContent.Schema.Value
			}
		}

		// Generate mock data based on schema
		mockData := generateMockData(schema)

		// Send response based on content type
		switch {
		case strings.Contains(contentType, "application/xml"):
			c.Header("Content-Type", "application/xml")
			xmlData, err := xml.Marshal(mockData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate XML"})
				return
			}
			c.String(status, string(xmlData))
		default:
			c.JSON(status, mockData)
		}
	}
}

func (m *MockServer) findResponse(op *openapi3.Operation) (int, *openapi3.Response) {
	for statusCode, responseRef := range op.Responses.Map() {
		status := parseStatusCode(statusCode)
		if status >= 200 && status < 300 {
			return status, responseRef.Value
		}
	}
	// Default to 200 if no successful response found
	if defaultResponse := op.Responses.Default(); defaultResponse != nil {
		return 200, defaultResponse.Value
	}
	return 200, nil
}
