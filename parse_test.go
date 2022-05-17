package searchable_test

import (
	"reflect"
	"testing"

	. "github.com/bsm/go-searchable"
)

func testParse(t *testing.T, s string, exp []Token) {
	t.Helper()

	if got := Parse(s); !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %v, but got %v", exp, got)
	}
}

func TestParse(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		testParse(t, "", []Token{})
	})
	t.Run("blank doubles", func(t *testing.T) {
		testParse(t, `""`, []Token{})
	})
	t.Run("-+", func(t *testing.T) {
		testParse(t, `-+`, []Token{
			{Term: "+", Negate: true},
		})
	})
	t.Run("simple words", func(t *testing.T) {
		testParse(t, "simple words", []Token{
			{Term: "simple"},
			{Term: "words"},
		})
	})
	t.Run("with spaces", func(t *testing.T) {
		testParse(t, " with   \t spaces\n", []Token{
			{Term: "with"},
			{Term: "spaces"},
		})
	})
	t.Run("with duplicates", func(t *testing.T) {
		testParse(t, "with with duplicates with", []Token{
			{Term: "with"},
			{Term: "duplicates"},
		})
	})
	t.Run("with full term", func(t *testing.T) {
		testParse(t, `with "full term"`, []Token{
			{Term: "with"},
			{Term: "full term"},
		})
	})
	t.Run("odd double quotes", func(t *testing.T) {
		testParse(t, `"""odd double quotes around"""`, []Token{
			{Term: "odd double quotes around"},
		})
	})
	t.Run("even double quotes", func(t *testing.T) {
		testParse(t, `""even double quotes around""`, []Token{
			{Term: "even double quotes around"},
		})
	})
	t.Run("with apostrophe", func(t *testing.T) {
		testParse(t, `with'apostrophe`, []Token{
			{Term: "with'apostrophe"},
		})
	})
	t.Run("with -minus", func(t *testing.T) {
		testParse(t, `with -minus`, []Token{
			{Term: "with"},
			{Term: "minus", Negate: true},
		})
	})
	t.Run("with +plus", func(t *testing.T) {
		testParse(t, `with +plus`, []Token{
			{Term: "with"},
			{Term: "plus"},
		})
	})
	t.Run("with-minus", func(t *testing.T) {
		testParse(t, `with-minus`, []Token{
			{Term: "with-minus"},
		})
	})
	t.Run("with+plus", func(t *testing.T) {
		testParse(t, `with+plus`, []Token{
			{Term: "with+plus"},
		})
	})
	t.Run("with minus before", func(t *testing.T) {
		testParse(t, `with -"minus before"`, []Token{
			{Term: "with"},
			{Term: "minus before", Negate: true},
		})
	})
	t.Run("with minus within", func(t *testing.T) {
		testParse(t, `with "-minus within"`, []Token{
			{Term: "with"},
			{Term: "-minus within"},
		})
	})
	t.Run("with plus before", func(t *testing.T) {
		testParse(t, `with +"plus before"`, []Token{
			{Term: "with"},
			{Term: "plus before"},
		})
	})
	t.Run("with plus within", func(t *testing.T) {
		testParse(t, `with "+plus within"`, []Token{
			{Term: "with"},
			{Term: "+plus within"},
		})
	})
	t.Run("+plus in other term", func(t *testing.T) {
		testParse(t, `+plus "in other term"`, []Token{
			{Term: "plus"},
			{Term: "in other term"},
		})
	})
	t.Run("with blank", func(t *testing.T) {
		testParse(t, `with ''`, []Token{
			{Term: "with"},
		})
	})
	t.Run("with blank doubles", func(t *testing.T) {
		testParse(t, `with ""`, []Token{
			{Term: "with"},
		})
	})
}
