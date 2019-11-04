package searchable

type conditions struct {
	SQL    []byte
	Args   []interface{}
	Len    int
	Negate bool
}

func (c *conditions) Append(sql, op string, val interface{}) {
	if c.Len == 0 {
		if c.Negate {
			c.SQL = append(c.SQL, '!')
		}
		c.SQL = append(c.SQL, '(')
	} else {
		c.SQL = append(c.SQL, " OR "...)
	}
	c.SQL = append(c.SQL, '(')
	c.SQL = append(c.SQL, sql...)
	c.SQL = append(c.SQL, " IS NOT NULL AND "...)
	c.SQL = append(c.SQL, sql...)
	c.SQL = append(c.SQL, ' ')
	c.SQL = append(c.SQL, op...)
	c.SQL = append(c.SQL, ' ', '?', ')')
	c.Args = append(c.Args, val)
	c.Len++
}

func (c *conditions) Done() {
	if c.Len != 0 {
		c.SQL = append(c.SQL, ')')
	}
}

// ToSql implements squirrel.Sqlizer.
func (c *conditions) ToSql() (string, []interface{}, error) {
	return string(c.SQL), c.Args, nil
}
