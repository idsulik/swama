package model

const (
	ParameterInPath   = "path"
	ParameterInQuery  = "query"
	ParameterInHeader = "header"
	ParameterInCookie = "cookie"
)

type Parameter struct {
	In          string
	Name        string
	Type        string
	Required    bool
	Description string
}

func NewParameter(in string, name string, typ string, required bool, description string) *Parameter {
	return &Parameter{
		In:          in,
		Name:        name,
		Type:        typ,
		Required:    required,
		Description: description,
	}
}
