package searchable

import (
	"regexp"
	"strings"
)

var tokenPattern = regexp.MustCompile(`[\-\+]?("+([^"]*)"+|[^\s]+)`)

// Parse parses search tokens from a plain string.
func Parse(s string) []Token {
	raw := tokenPattern.FindAllString(s, -1)
	res := make([]Token, 0, len(raw))

	for _, term := range raw {
		sign := ""
		if strings.HasPrefix(term, "-") || strings.HasPrefix(term, "+") {
			sign, term = string(term[0]), term[1:]
		}

		term = strings.Trim(term, `"`)
		if term == "" || term == `''` {
			continue
		}

		var val Token
		if sign == "-" {
			val = Token{Term: term, Negate: true}
		} else if sign == "+" {
			val = Token{Term: term, Negate: false}
		} else {
			val = Token{Term: sign + term, Negate: false}
		}

		var exists bool
		for _, v := range res {
			if v == val {
				exists = true
				break
			}
		}

		if !exists {
			res = append(res, val)
		}
	}
	return res
}
