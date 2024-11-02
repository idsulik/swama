package mockserver

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/getkin/kin-openapi/openapi3"
)

func generateMockData(schema *openapi3.Schema) interface{} {
	if schema == nil {
		return nil
	}

	switch strings.Join(schema.Type.Slice(), "") {
	case "object":
		obj := make(map[string]interface{})
		for propName, propSchema := range schema.Properties {
			obj[propName] = generateMockData(propSchema.Value)
		}
		return obj
	case "array":
		arr := make([]interface{}, 0)
		count := gofakeit.Number(1, 5)
		for i := 0; i < count; i++ {
			arr = append(arr, generateMockData(schema.Items.Value))
		}
		return arr
	case "string":
		switch schema.Format {
		case "date-time":
			return gofakeit.Date().Format(time.RFC3339)
		case "date":
			return gofakeit.Date().Format("2006-01-02")
		case "email":
			return gofakeit.Email()
		case "uuid":
			return gofakeit.UUID()
		default:
			if schema.Example != nil {
				return schema.Example
			}
			return gofakeit.Sentence(3)
		}
	case "number", "integer":
		if schema.Example != nil {
			return schema.Example
		}
		return gofakeit.Number(1, 1000)
	case "boolean":
		return gofakeit.Bool()
	default:
		return nil
	}
}

func parseAcceptHeader(header string) []string {
	if header == "" {
		return []string{"application/json"}
	}

	types := strings.Split(header, ",")
	for i, t := range types {
		if idx := strings.Index(t, ";"); idx != -1 {
			types[i] = strings.TrimSpace(t[:idx])
		} else {
			types[i] = strings.TrimSpace(t)
		}
	}
	return types
}

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
