package searchable_test

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/bsm/go-searchable"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	var subject = searchable.Builder{
		{SQL: "users.name"},
		{SQL: "users.age", Type: searchable.TypeInt},
		{SQL: "users.code", Exact: true},
	}

	It("should search strings", func() {
		search := subject.SearchStrings([]string{"alice", "45", "-admin"})
		Expect(search).To(BeAssignableToTypeOf(squirrel.And{}))

		sql, args, err := search.ToSql()
		Expect(err).NotTo(HaveOccurred())

		Expect(sql).To(Equal(`(` +
			`((users.name IS NOT NULL AND users.name LIKE ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			` AND ` +
			`((users.name IS NOT NULL AND users.name LIKE ?) OR (users.age IS NOT NULL AND users.age = ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			` AND ` +
			`((users.name IS NOT NULL AND users.name LIKE ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			`)`,
		))
		Expect(args).To(Equal([]interface{}{
			"%alice%", "alice",
			"%45%", int64(45), "45",
			"%-admin%", "-admin",
		}))
	})

	It("should search parsed terms and negate where necessary", func() {
		search := subject.Search([]searchable.Token{{Term: "alice"}, {Term: "45"}, {Term: "admin", Negate: true}})
		Expect(search).To(BeAssignableToTypeOf(squirrel.And{}))

		sql, args, err := search.ToSql()
		Expect(err).NotTo(HaveOccurred())
		Expect(sql).To(Equal(`(` +
			`((users.name IS NOT NULL AND users.name LIKE ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			` AND ` +
			`((users.name IS NOT NULL AND users.name LIKE ?) OR (users.age IS NOT NULL AND users.age = ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			` AND ` +
			`!((users.name IS NOT NULL AND users.name LIKE ?) OR (users.code IS NOT NULL AND users.code = ?))` +
			`)`,
		))
		Expect(args).To(Equal([]interface{}{
			"%alice%", "alice",
			"%45%", int64(45), "45",
			"%admin%", "admin",
		}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "searchable")
}
