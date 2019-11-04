package searchable

import (
	"regexp"
	"strings"
)

var tokenPattern = regexp.MustCompile(`[\-\+]?("+([^"]*)"+|[^\s]+)`)

// ParseTokens parses search tokens from a plain string
func ParseTokens(s string) []string {
	raw := tokenPattern.FindAllString(s, -1)
	res := raw[:0]
	for _, t := range raw {
		p := ""
		if strings.HasPrefix(t, "-") || strings.HasPrefix(t, "+") {
			p, t = string(t[0]), t[1:len(t)]
		}

		x := strings.Trim(t, `"`)
		if x != "" {
			res = append(res, p+x)
		}
	}
	return res
}
