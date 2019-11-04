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

// Clause defines a signgle field clause that will be used for
// matching. This can be a simple column name but also contain
// complex expressions.
type Clause struct {
	SQL   string
	Type  Type
	Exact bool // support exact matches only (for string types)
}

// Builder builds a search expression.
type Builder []Clause

// Search converts a list of search tokens into a WHERE clause.
func (b Builder) Search(tokens []string) squirrel.Sqlizer {
	var conj squirrel.And
	for _, token := range tokens {
		if part := b.buildPart(token); part != nil {
			conj = append(conj, part)
		}
	}
	return conj
}

func (b Builder) buildPart(token string) squirrel.Sqlizer {
	var pattern string

	term := strings.TrimLeft(token, "+-")
	if term == "" {
		return nil
	}

	cc := &conditions{Negate: strings.HasPrefix(token, "-")}
	for _, clause := range b {
		switch clause.Type {
		case TypeString:
			if clause.Exact {
				cc.Append(clause.SQL, "=", term)
			} else {
				if pattern == "" {
					pattern = `%` + likeEscaper.Replace(term) + `%`
				}
				cc.Append(clause.SQL, "LIKE", pattern)
			}
		case TypeInt:
			if num, err := strconv.ParseInt(term, 10, 64); err == nil {
				cc.Append(clause.SQL, "=", num)
			}
		}
	}
	cc.Done()
	return cc
}
