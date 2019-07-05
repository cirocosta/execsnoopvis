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

			Context("without branching", func() {

				BeforeEach(func() {
					node1 := &parser.Node{Pid: 1, Command: "foo"}
					node2 := &parser.Node{Pid: 2, Command: "bar", Parent: node1}
					node1.Children = []*parser.Node{node2}

					roots = []*parser.Node{node1}
				})

				It("succeeds", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("prints a dot-formatted graph", func() {
					Expect(content).To(Equal(`digraph G {
	foo_1->bar_2;
	bar_2;
	foo_1;

}
`))

				})

			})

			Context("with branching", func() {

				BeforeEach(func() {
					node1 := &parser.Node{Pid: 1, Command: "foo"}
					node2 := &parser.Node{Pid: 2, Command: "bar", Parent: node1}
					node3 := &parser.Node{Pid: 3, Command: "caz", Parent: node1}
					node1.Children = []*parser.Node{node2, node3}

					roots = []*parser.Node{node1}
				})

				It("succeeds", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("prints a dot-formatted graph", func() {
					Expect(content).To(Equal(`digraph G {
	foo_1->bar_2;
	foo_1->caz_3;
	bar_2;
	caz_3;
	foo_1;

}
`))

				})

			})

			Context("with multiple roots", func() {

				BeforeEach(func() {
					node1 := &parser.Node{Pid: 1, Command: "bar"}
					node2 := &parser.Node{Pid: 2, Command: "caz"}

					roots = []*parser.Node{node1, node2}
				})

				It("succeeds", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("prints a dot-formatted graph", func() {
					Expect(content).To(Equal(`digraph G {
	bar_1;
	caz_2;

}
`))

				})

			})

		})

	})

})
