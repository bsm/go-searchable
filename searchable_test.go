package searchable_test

import (
	"reflect"
	"testing"

	. "github.com/bsm/go-searchable"
)

func TestBuilder(t *testing.T) {
	b := &Builder{
		Fields: []Field{
			{SQL: "users.name"},
			{SQL: "users.age", Type: FieldInt},
			{SQL: "users.code", Match: MatchExact},
		},
		Placeholder: Question,
	}

	t.Run("Search", func(t *testing.T) {
		tokens := Parse("alice 45 -admin")
		clause := b.Search(tokens)
		sql, args, err := clause.ToSql()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		} else if exp := `` +
			`((users.name IS NOT NULL AND users.name ILIKE ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			` AND ` +
			`((users.name IS NOT NULL AND users.name ILIKE ?) OR (users.age IS NOT NULL AND users.age = ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			` AND ` +
			`!((users.name IS NOT NULL AND users.name ILIKE ?) OR (users.code IS NOT NULL AND users.code = ?))`; exp != sql {
			t.Errorf("expected:\n\t%v,\nbut got:\n\t%v", exp, sql)
		}

		if exp := []interface{}{
			"%alice%", "alice",
			"%45%", int64(45), "45",
			"%admin%", "admin",
		}; !reflect.DeepEqual(exp, args) {
			t.Errorf("expected %v, but got %v", exp, args)
		}
	})
}
