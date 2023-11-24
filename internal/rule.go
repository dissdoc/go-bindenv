package internal

import "regexp"

var rule *regexp.Regexp

func GetRule() *regexp.Regexp {
	if rule == nil {
		rule = regexp.MustCompile(`^\$\{(.*?)\}`)
	}

	return rule
}
