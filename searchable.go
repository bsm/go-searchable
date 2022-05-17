package searchable

// FieldType reflects on the type of the field.
type FieldType uint8

// FieldType enum.
const (
	FieldString FieldType = iota
	FieldInt
)

// MatchType defines the type of the match. Strings only.
type MatchType uint8

// MatchType enum.
const (
	MatchContains MatchType = iota
	MatchExact
	MatchPrefix
)

// PlaceholderFormat represents a specific format.
type PlaceholderFormat uint8

const (
	// Dollar placeholder e.g. $1, $2, $3, commonly used by PostgreSQL.
	Dollar PlaceholderFormat = iota
	// Question placeholder (e.g. ?, ?, ?), commonly used by MySQL.
	Question
	// Colon placeholder (e.g. :1, :2, :3), commonly used by Oracle.
	Colon
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
type Field struct {
	SQL   string
	Type  FieldType
	Match MatchType
}

// Builder builds a search expression.
type Builder struct {
	Fields      []Field
	Placeholder PlaceholderFormat
}

// Search converts a list of parsed search tokens into a WHERE Clause.
func (b *Builder) Search(tokens []Token) *Clause {
	return &Clause{builder: b, tokens: tokens}
}
