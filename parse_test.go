package searchable_test

import (
	"github.com/bsm/go-searchable"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("ParseTokens",
	func(s string, x []searchable.Token) {
		Expect(searchable.ParseTokens(s)).To(Equal(x))
	},
	Entry("empty", "", []searchable.Token{}),
	Entry("blank doubles", `""`, []searchable.Token{}),
	Entry("-+", `-+`, []searchable.Token{
		{Term: "+", Negate: true},
	}),
	Entry("simple words", "simple words", []searchable.Token{
		{Term: "simple"},
		{Term: "words"},
	}),
	Entry("with spaces", " with   \t spaces\n", []searchable.Token{
		{Term: "with"},
		{Term: "spaces"},
	}),
	Entry("with duplicates", "with with duplicates with", []searchable.Token{
		{Term: "with"},
		{Term: "duplicates"},
	}),
	Entry("with full term", `with "full term"`, []searchable.Token{
		{Term: "with"},
		{Term: "full term"},
	}),
	Entry("odd double quotes", `"""odd double quotes around"""`, []searchable.Token{
		{Term: "odd double quotes around"},
	}),
	Entry("even double quotes", `""even double quotes around""`, []searchable.Token{
		{Term: "even double quotes around"},
	}),
	Entry("with apostrophe", `with'apostrophe`, []searchable.Token{
		{Term: "with'apostrophe"},
	}),
	Entry("with -minus", `with -minus`, []searchable.Token{
		{Term: "with"},
		{Term: "minus", Negate: true},
	}),
	Entry("with +plus", `with +plus`, []searchable.Token{
		{Term: "with"},
		{Term: "plus"},
	}),
	Entry("with-minus", `with-minus`, []searchable.Token{
		{Term: "with-minus"},
	}),
	Entry("with+plus", `with+plus`, []searchable.Token{
		{Term: "with+plus"},
	}),
	Entry("with minus before", `with -"minus before"`, []searchable.Token{
		{Term: "with"},
		{Term: "minus before", Negate: true},
	}),
	Entry("with minus within", `with "-minus within"`, []searchable.Token{
		{Term: "with"},
		{Term: "-minus within"},
	}),
	Entry("with plus before", `with +"plus before"`, []searchable.Token{
		{Term: "with"},
		{Term: "plus before"},
	}),
	Entry("with plus within", `with "+plus within"`, []searchable.Token{
		{Term: "with"},
		{Term: "+plus within"},
	}),
	Entry("+plus in other term", `+plus "in other term"`, []searchable.Token{
		{Term: "plus"},
		{Term: "in other term"},
	}),
	Entry("with blank", `with ''`, []searchable.Token{
		{Term: "with"},
	}),
	Entry("with blank doubles", `with ""`, []searchable.Token{
		{Term: "with"},
	}),
)
