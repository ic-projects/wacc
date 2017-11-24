package utils

import (
	"bytes"
	"fmt"
	"strings"
)

// Indent is a function to indent when printing the AST. Given the string s, it
// indents the string with all previous indents plus the new Indent (sep).
func Indent(s string, sep string) string {
	var buf bytes.Buffer
	for _, line := range strings.Split(s, "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%s%s\n", sep, line))
		}
	}
	return buf.String()
}
