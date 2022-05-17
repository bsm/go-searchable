package searchable_test

import (
	"fmt"

	"github.com/bsm/go-searchable"
)

func ExampleBuilder_Search() {
	builder := &searchable.Builder{
		Fields: []searchable.Field{
			{SQL: "users.name"},
			{SQL: "users.code", Match: searchable.MatchExact},
		},
		Placeholder: searchable.Dollar,
	}

	tokens := searchable.Parse("alice")
	search := builder.Search(tokens)
	sql, args, _ := search.ToSql()
	fmt.Printf("%v\n", sql)
	fmt.Printf("%v\n", args)

	// Output:
	// ((users.name IS NOT NULL AND users.name ILIKE $1) OR (users.code IS NOT NULL AND users.code = $2))
	// [%alice% alice]
}
