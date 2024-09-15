package openapi

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

var defaultFiles = []string{
	"swagger.yaml",
	"swagger.yml",
	"swagger.json",
	"openapi.yaml",
	"openapi.yml",
	"openapi.json",
}

// ParseSwaggerFile parses a Swagger file (YAML/JSON) using kin-openapi, with support for context cancellation.
func ParseSwaggerFile(ctx context.Context, filePath string) (*openapi3.T, error) {
	loader := openapi3.NewLoader()

	// Use the context for any long-running operations or network requests.
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("operation canceled")
	default:
		doc, err := loader.LoadFromFile(filePath)
		if err != nil {
			return nil, err
		}

		err = doc.Validate(ctx)
		if err != nil {
			return nil, err
		}

		return doc, nil
	}
}

// LocateSwaggerFile tries to find the Swagger file in the current directory.
func LocateSwaggerFile() string {
	for _, file := range defaultFiles {
		if _, err := os.Stat(file); err == nil {
			log.Printf("Using Swagger file: %s\n", file)
			return file
		}
	}

	log.Println("Swagger file not found in the current directory.")
	return ""
}

func ViewEndpointDetails(swagger *openapi3.T, endpoint string) (string, error) {
	path := swagger.Paths.Find(endpoint)
	if path == nil {
		return "", fmt.Errorf("endpoint %s not found", endpoint)
	}

	var details string
	for _, operation := range path.Operations() {
		details += fmt.Sprintf("[%s] %s\n", operation.OperationID, endpoint)
		if operation.Summary != "" {
			details += fmt.Sprintf("  - %s\n", operation.Summary)
		}
	}

	return details, nil
}
