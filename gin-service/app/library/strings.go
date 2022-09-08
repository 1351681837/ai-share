package library

import (
	"strings"
	"unicode"
)

/**
*驼峰转下划线
*/
func CamelCaseToUdnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

/**
 *下划线转驼峰
 */
func UderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}