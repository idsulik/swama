package util

import (
	"strings"

	"github.com/olekukonko/tablewriter"
)

func GetMethodColor(method string) int {
	method = strings.ToUpper(method)
	methodColor := tablewriter.FgHiWhiteColor
	if method == "GET" {
		methodColor = tablewriter.FgHiBlueColor
	} else if method == "POST" {
		methodColor = tablewriter.FgHiGreenColor
	} else if method == "PUT" {
		methodColor = tablewriter.FgHiYellowColor
	} else if method == "DELETE" {
		methodColor = tablewriter.FgHiRedColor
	}

	return methodColor
}
