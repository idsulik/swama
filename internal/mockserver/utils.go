package mockserver

import (
	"encoding/xml"
	"fmt"
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

// mapToXML converts a map[string]interface{} to XML structure
func mapToXML(data interface{}, schema *openapi3.Schema, name string) *XMLNode {
	if data == nil || schema == nil {
		return nil
	}

	// Get XML name from schema or use provided name
	xmlName := name
	if schema.XML != nil && schema.XML.Name != "" {
		xmlName = schema.XML.Name
	}

	node := &XMLNode{
		XMLName: xml.Name{Local: xmlName},
	}

	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			propSchema := schema.Properties[key].Value
			if propSchema == nil {
				continue
			}

			// Handle properties marked as XML attributes
			if propSchema.XML != nil && propSchema.XML.Attribute {
				node.Attrs = append(
					node.Attrs, xml.Attr{
						Name:  xml.Name{Local: key},
						Value: fmt.Sprintf("%v", value),
					},
				)
				continue
			}

			// Get property name from schema or use key
			propName := key
			if propSchema.XML != nil && propSchema.XML.Name != "" {
				propName = propSchema.XML.Name
			}

			childNode := mapToXML(value, propSchema, propName)
			if childNode != nil {
				node.Children = append(node.Children, childNode)
			}
		}

	case []interface{}:
		// Handle array wrapping if specified in schema
		if schema.XML != nil && schema.XML.Wrapped {
			// For wrapped arrays, return the current node and append items as children
			for _, item := range v {
				childNode := mapToXML(item, schema.Items.Value, name)
				if childNode != nil {
					node.Children = append(node.Children, childNode)
				}
			}
			return node
		} else {
			// For unwrapped arrays, return an array of nodes
			nodes := make([]*XMLNode, 0)
			for _, item := range v {
				childNode := mapToXML(item, schema.Items.Value, name)
				if childNode != nil {
					nodes = append(nodes, childNode)
				}
			}
			// If this is the root node, wrap it
			if name != "" {
				node.Children = nodes
				return node
			}
			return &XMLNode{
				XMLName:  xml.Name{Local: "array"},
				Children: nodes,
			}
		}

	default:
		// Handle primitive values
		node.Value = fmt.Sprintf("%v", v)
	}

	return node
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

// convertPathToGinFormat Convert OpenAPI path params ({param}) to Gin format (:param)
func convertPathToGinFormat(path string) string {
	path = strings.ReplaceAll(path, "{", ":")
	path = strings.ReplaceAll(path, "}", "")

	return path
}

// convertGinPathToOpenAPI Convert Gin path params (:param) back to OpenAPI format ({param})
func convertGinPathToOpenAPI(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			parts[i] = "{" + strings.TrimPrefix(part, ":") + "}"
		}
	}
	return strings.Join(parts, "/")
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
