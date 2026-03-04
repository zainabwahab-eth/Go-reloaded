package operations

import (
	"regexp"
	"strings"
)

func FixQuotes(text string) string {
	re := regexp.MustCompile(`\s*'\s*(.*?)\s*'\s*`)
	result := re.ReplaceAllString(text, " '$1' ")
	return strings.TrimSpace(result)
}
