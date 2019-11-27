package searchable_test

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/bsm/go-searchable"
)

func ExampleBuilder_SearchStrings() {
	builder := searchable.Builder{
		{SQL: "users.name"},
		{SQL: "users.age", Type: searchable.TypeInt},
		{SQL: "users.code", Exact: true},
	}
	search := builder.SearchStrings([]string{"alice", "45", "-admin"})
	users := squirrel.Select("*").From("users").Where(search)
	sql, args, _ := users.ToSql()
	fmt.Printf("%v\n", sql)
	fmt.Printf("%v\n", args)

	// Output:
	// SELECT * FROM users WHERE (((users.name IS NOT NULL AND users.name LIKE ?) OR (users.code IS NOT NULL AND users.code = ?)) AND ((users.name IS NOT NULL AND users.name LIKE ?) OR (users.age IS NOT NULL AND users.age = ?) OR (users.code IS NOT NULL AND users.code = ?)) AND ((users.name IS NOT NULL AND users.name LIKE ?) OR (users.code IS NOT NULL AND users.code = ?)))
	// [%alice% alice %45% 45 45 %-admin% -admin]
}
