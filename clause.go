package searchable

import (
	"strconv"
	"strings"
	"unsafe"
)

var likeEscaper = strings.NewReplacer(`%`, `\%`, `_`, `\_`)

// Clause contains the WHERE clause.
type Clause struct {
	builder *Builder
	tokens  []Token
}

// ToSql implements github.com/Masterminds/squirrel.Sqlizer interface.
func (c *Clause) ToSql() (string, []interface{}, error) {
	var (
		sql  []byte
		args []interface{}
	)

	for _, t := range c.tokens {
		if t.Term == "" {
			continue
		}

		firstField := true
		for _, f := range c.builder.Fields {
			switch f.Type {
			case FieldString:
				switch f.Match {
				case MatchExact:
					sql = nextField(sql, firstField, t)
					sql = c.appendCondition(sql, f.SQL, "=", len(args))
					args = append(args, t.Term)
					firstField = false
				case MatchPrefix:
					sql = nextField(sql, firstField, t)
					sql = c.appendCondition(sql, f.SQL, "ILIKE", len(args))
					args = append(args, likeEscaper.Replace(t.Term)+`%`)
					firstField = false
				default:
					sql = nextField(sql, firstField, t)
					sql = c.appendCondition(sql, f.SQL, "ILIKE", len(args))
					args = append(args, `%`+likeEscaper.Replace(t.Term)+`%`)
					firstField = false
				}
			case FieldInt:
				if num, err := strconv.ParseInt(t.Term, 10, 64); err == nil {
					sql = nextField(sql, firstField, t)
					sql = c.appendCondition(sql, f.SQL, "=", len(args))
					args = append(args, num)
					firstField = false
				}
			}
		}

		if !firstField {
			sql = append(sql, ')')
		}
	}

	if len(sql) == 0 {
		return "TRUE", nil, nil
	}
	return *(*string)(unsafe.Pointer(&sql)), args, nil
}

func (c *Clause) appendPlacholder(sql []byte, numArgs int) []byte {
	switch c.builder.Placeholder {
	case Question:
		sql = append(sql, '?')
	case Colon:
		sql = append(sql, '?')
		sql = strconv.AppendInt(sql, int64(numArgs)+1, 10)
	default:
		sql = append(sql, '$')
		sql = strconv.AppendInt(sql, int64(numArgs)+1, 10)
	}
	return sql
}

func (c *Clause) appendCondition(sql []byte, field, op string, numArgs int) []byte {
	sql = append(sql, '(')
	sql = append(sql, field...)
	sql = append(sql, " IS NOT NULL AND "...)
	sql = append(sql, field...)
	sql = append(sql, ' ')
	sql = append(sql, op...)
	sql = append(sql, ' ')
	sql = c.appendPlacholder(sql, numArgs)
	sql = append(sql, ')')
	return sql
}

func nextField(sql []byte, firstField bool, t Token) []byte {
	if firstField {
		if len(sql) != 0 {
			sql = append(sql, ' ', 'A', 'N', 'D', ' ')
		}
		if t.Negate {
			sql = append(sql, '!')
		}
		sql = append(sql, '(')
	} else {
		sql = append(sql, ' ', 'O', 'R', ' ')
	}
	return sql
}
