package osUtils

import (
	"log"
	"regexp"
)

func GetRegexMatches(pattern string, input string) []string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error compiling regex pattern: (", pattern, ") ", err)
	}
	return re.FindAllString(input, -1)
}