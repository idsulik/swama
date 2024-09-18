package swagger

import (
	"context"
	"log"
	"net/url"
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

// LoadSwaggerFile loads the Swagger/OpenAPI file into a parsed document.
func LoadSwaggerFile(ctx context.Context, filepath string) (*openapi3.T, error) {
	swaggerLoader := &openapi3.Loader{
		Context:               ctx,
		IsExternalRefsAllowed: true,
	}

	url, err := url.Parse(filepath)
	if err != nil {
		return nil, err
	}

	doc, err := swaggerLoader.LoadFromURI(url)
	if err != nil {
		return nil, err
	}

	return doc, nil
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
