package main

import (
	"strings"
)

func headerBlock(rolePath string)(string) {
	var header []string
	var headerBlockText string
	header = strings.Split(rolePath, "\\")
	headerBlockText = strings.Replace(header[len(header)-1], "-", " ", -1)
	headerBlockText = "# " + strings.Title(headerBlockText)

	return headerBlockText
}