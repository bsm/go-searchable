package searchable_test

import (
	"github.com/bsm/go-searchable"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("ParseTokens",
	func(s string, x []string) {
		Expect(searchable.ParseTokens(s)).To(Equal(x))
	},
	Entry("empty", "", nil),
	Entry("simple words", "simple words", []string{"simple", "words"}),
	Entry("with spaces", " with   \t spaces\n", []string{"with", "spaces"}),
	Entry("with full term", `with "full term"`, []string{"with", "full term"}),
	Entry("odd double quotes", `"""odd double quotes around"""`, []string{"odd double quotes around"}),
	Entry("even double quotes", `""even double quotes around""`, []string{"even double quotes around"}),
	Entry("with minus before", `with -"minus before"`, []string{"with", "-minus before"}),
	Entry("with minus within", `with "-minus within"`, []string{"with", "-minus within"}),
	Entry("with plus before", `with +"plus before"`, []string{"with", "+plus before"}),
	Entry("with plus within", `with "+plus within"`, []string{"with", "+plus within"}),
	Entry("with blank", `with ""`, []string{"with"}),
)
