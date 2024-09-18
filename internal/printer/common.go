package printer

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func getProperties(value *openapi3.Schema) []string {
	properties := make([]string, 0)

	if value.Properties != nil {
		sortedPropertiesName := slices.Sorted(maps.Keys(value.Properties))
		for _, propertyName := range sortedPropertiesName {
			prop := value.Properties[propertyName]
			propertyName = enrichPropertyName(propertyName, prop)
			properties = append(properties, propertyName)
		}
	}

	return properties
}

func enrichPropertyName(propertyName string, prop *openapi3.SchemaRef) string {
	if prop == nil || prop.Value == nil {
		return propertyName
	}

	if prop.Value.Type != nil {
		if prop.Value.Format == "" {
			propertyName += fmt.Sprintf(" (%s)", strings.Join(prop.Value.Type.Slice(), ", "))
		} else {
			propertyName += fmt.Sprintf(" (%s: %s)", strings.Join(prop.Value.Type.Slice(), ", "), prop.Value.Format)
		}

		if prop.Value.Properties != nil {
			properties := getProperties(prop.Value)
			propertyName += fmt.Sprintf(" {%s}", strings.Join(properties, ","))
		}
	}

	return propertyName
}
