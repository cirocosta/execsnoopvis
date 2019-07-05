package parser_test

import (
	"github.com/cirocosta/execsnoopvis/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vis", func() {

	Describe("Render", func() {

		var (
			err     error
			roots   []*parser.Node
			content string
		)

		JustBeforeEach(func() {
			content, err = parser.Render(roots)
		})

		Context("with no nodes", func() {

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})

		})

		Context("having a non-empty node tree", func() {

			BeforeEach(func() {
				node1 := &parser.Node{}
				node2 := &parser.Node{Parent: node1}
				node1.Children = []*parser.Node{node2}

				roots = []*parser.Node{node1}
			})

			It("succeeds", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("prints a dot-formatted graph", func() {
				Expect(content).To(Equal(`digraph G {

}
`))
			})

		})

	})

})
