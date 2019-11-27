# Go Searchable

[![Build Status](https://travis-ci.org/bsm/go-searchable.png?branch=master)](https://travis-ci.org/bsm/go-searchable)
[![GoDoc](https://godoc.org/github.com/bsm/go-searchable?status.png)](http://godoc.org/github.com/bsm/go-searchable)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsm/go-searchable)](https://goreportcard.com/report/github.com/bsm/go-searchable)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Simple search query builder on top of [squirrel](https://github.com/Masterminds/squirrel).

## Usage

```go
import (
  "fmt"

  "github.com/bsm/go-searchable"
  "github.com/Masterminds/squirrel"
)

var builder = searchable.Builder{
  {SQL: "users.name"},
  {SQL: "users.age", Type: searchable.TypeInt},
  {SQL: "users.code", Exact: true},
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
  // => SELECT * FROM users WHERE (((users.name IS NOT NULL AND users.name LIKE ?) OR (users.code IS NOT NULL AND users.code = ?)) AND ((users.name IS NOT NULL AND users.name LIKE ?) OR ...
  fmt.Println(args)
  // => [%alice% alice %45% 45 45 %admin% admin]
}
```
