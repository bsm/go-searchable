package searchable

import (
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"
)

var likeEscaper = strings.NewReplacer(`%`, `\%`, `_`, `\_`)

// Type reflects on the type of the clause.
type Type uint8

// Type enum.
const (
	TypeString Type = iota
	TypeInt
)

// Token defines an individual term that may or may not be
// negated from the search.
type Token struct {
	Term   string
	Negate bool
}

// Clause defines a single field clause that will be used for
// matching. This can be a simple column name but also contain
// complex expressions.
type Clause struct {
	SQL   string
	Type  Type
	Exact bool // support exact matches only (for string types)
}

// Builder builds a search expression.
type Builder []Clause

// Search converts a list of parsed search tokens into a WHERE clause.
func (b Builder) Search(tokens []Token) squirrel.Sqlizer {
	var conj squirrel.And
	for _, token := range tokens {
		if part := b.buildPart(token); part != nil {
			conj = append(conj, part)
		}
	}
	return conj
}

// SearchStrings converts a list of string search tokens into a WHERE clause without excluding terms preceded by "-".
func (b Builder) SearchStrings(terms []string) squirrel.Sqlizer {
	// convert strings to []Token with Negate: false then just pass in to search overall method
	tokens := make([]Token, 0, len(terms))
	for _, term := range terms {
		tokens = append(tokens, Token{Term: term})
	}
	return b.Search(tokens)
}

func (b Builder) buildPart(token Token) squirrel.Sqlizer {
	var pattern string

	if token.Term == "" {
		return nil
	}

	cc := &conditions{Negate: token.Negate}
	for _, clause := range b {
		switch clause.Type {
		case TypeString:
			if clause.Exact {
				cc.Append(clause.SQL, "=", token.Term)
			} else {
				if pattern == "" {
					pattern = `%` + likeEscaper.Replace(token.Term) + `%`
				}
				cc.Append(clause.SQL, "LIKE", pattern)
			}
		case TypeInt:
			if num, err := strconv.ParseInt(token.Term, 10, 64); err == nil {
				cc.Append(clause.SQL, "=", num)
			}
		}
	}
	cc.Done()
	return cc
}
