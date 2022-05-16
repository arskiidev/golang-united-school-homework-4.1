package reverse_string

import (
	"strings"
)

func ReverseString(input string) (output string) {
	var l = []string{}
	var sb strings.Builder
	for _, v := range input {
		l = append(l, string(v))
	}
	for i := 1; i <= len(l); i++ {
		sb.WriteString(l[len(l)-i])
	}
	output = sb.String()
	return output
}
