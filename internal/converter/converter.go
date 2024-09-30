package converter

import (
	"errors"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/idsulik/swama/internal/model"
	"github.com/manifoldco/promptui"
)

const (
	CurlType  = "curl"
	FetchType = "fetch"
)

var (
	ErrInvalidConvertType = errors.New(fmt.Sprintf("invalid convert type. Must be %q or %q", CurlType, FetchType))
)

type Converter interface {
	ConvertEndpoint(method string, endpoint string, _ *model.Operation) string
}

func NewConverter(convertType string) (Converter, error) {
	switch convertType {
	case CurlType:
		return NewCurlConverter(), nil
	case FetchType:
		return NewFetchConverter(), nil
	default:
		return nil, ErrInvalidConvertType
	}
}

func askForValue(param *model.Parameter) string {
	var paramValue string
	fmt.Printf("Enter value for parameter %q: ", param.Name)
	_, _ = fmt.Scanln(&paramValue)
	if paramValue == "" && param.Required {
		fmt.Printf("parameter %q is required\n", param.Name)
		return askForValue(param)
	}

	return paramValue
}

func askForContentType(content openapi3.Content) string {
	if len(content) == 0 {
		return ""
	}

	contentTypes := make([]string, 0, len(content))
	for contentType := range content {
		contentTypes = append(contentTypes, contentType)
	}

	if len(contentTypes) == 1 {
		return contentTypes[0]
	}

	prompt := promptui.Select{
		Label: "Content Type",
		Items: contentTypes,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return askForContentType(content)
	}

	return result
}
