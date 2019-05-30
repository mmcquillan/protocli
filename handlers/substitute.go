package handlers

import (
	"regexp"
	"strings"
)

func Substitute(value string, tokens map[string]string) string {
	if match, hits := findVars(value); match {
		for _, hit := range hits {
			value = strings.Replace(value, hit, tokens[strip(hit)], -1)
		}
	}
	return value
}

func findVars(value string) (match bool, tokens []string) {
	match = false
	re := regexp.MustCompile("\\${([A-Za-z0-9:*_\\|\\-\\.\\?]+)}")
	tokens = re.FindAllString(strings.Replace(value, "$${", "X{", -1), -1)
	if len(tokens) > 0 {
		match = true
	}
	return match, tokens
}

func strip(value string) (stripped string) {
	stripped = strings.Replace(value, "${", "", -1)
	stripped = strings.Replace(stripped, "}", "", -1)
	return stripped
}
