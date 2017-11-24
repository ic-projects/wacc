package utils

import (
	"bytes"
	"strings"
	"fmt"
)

// Indent is a function to Indent when printing the AST, given the string s, it indents
// it with all previous indents plus the new Indent (sep)
func Indent(s string, sep string) string {
	var buf bytes.Buffer
	for _, line := range strings.Split(s, "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%s%s\n", sep, line))
		}
	}
	return buf.String()
}