# Go Searchable

[![Go Reference](https://pkg.go.dev/badge/github.com/bsm/go-searchable.svg)](https://pkg.go.dev/github.com/bsm/go-searchable)[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Simple search query builder, compatible with [squirrel](https://github.com/Masterminds/squirrel).

## Usage

```go
import (
  "fmt"

  "github.com/bsm/go-searchable"
  "github.com/Masterminds/squirrel"
)

var builder = searchable.Builder{
  Fields: []searchable.Field{
    {SQL: "users.name"},
    {SQL: "users.age", Type: searchable.FieldInt},
    {SQL: "users.code", Match: searchable.MatchExact},
  },
}

func main() {
  search := builder.Search([]searchable.Token{
    {Term: "alice"},
    {Term: "45"},
    {Term: "admin", Negate: true},
  })
  users := squirrel.Select("*").From("users").Where(search)
  sql, args, _ := users.ToSql()

  fmt.Println(sql)
  // => SELECT * FROM users WHERE (((users.name IS NOT NULL AND users.name LIKE $1) OR (users.code IS NOT NULL AND users.code = $2)) AND ((users.name IS NOT NULL AND users.name LIKE $3) OR ...
  fmt.Println(args)
  // => [%alice% alice %45% 45 45 %admin% admin]
}
```
