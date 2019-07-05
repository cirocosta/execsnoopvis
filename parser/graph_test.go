package parser_test

import (
	"github.com/cirocosta/execsnoopvis/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("graph", func() {

	var (
		nodes []*parser.Node
		roots []*parser.Node
	)

	Describe("FindRoots", func() {

		JustBeforeEach(func() {
			roots = parser.FindRoots(nodes)
		})

		Context("having no relationships", func() {

			BeforeEach(func() {
				nodes = []*parser.Node{
					{Parent: nil},
					{Parent: nil},
				}
			})

			It("all nodes are roots", func() {
				Expect(roots).To(ConsistOf(nodes))
			})

		})

		Context("having relationships", func() {

			BeforeEach(func() {
				nodes = []*parser.Node{{Parent: nil}}
				nodes = append(nodes, &parser.Node{
					Parent: nodes[0],
				})
			})

			It("returns the nodes that are roots", func() {
				Expect(roots).To(ConsistOf([]*parser.Node{
					nodes[0],
				}))
			})

		})

	})

	Describe("PopulateNodes", func() {

		JustBeforeEach(func() {
			parser.PopulateNodes(nodes)
		})

		Context("having nodes with relationships", func() {

			BeforeEach(func() {
				nodes = []*parser.Node{
					{
						Pid:  10,
						Ppid: 1,
					},
					{
						Pid:  11,
						Ppid: 10,
					},
				}
			})

			It("sets children", func() {
				Expect(nodes[0].Children).To(ConsistOf([]*parser.Node{
					nodes[1],
				}))
			})

			It("sets parent", func() {
				Expect(nodes[1].Parent).To(Equal(nodes[0]))
			})

		})

		Context("having nodes without any relationship", func() {

			BeforeEach(func() {
				nodes = []*parser.Node{
					{
						Pid:  10,
						Ppid: 1,
					},
					{
						Pid:  20,
						Ppid: 2,
					},
				}
			})

			It("doesn't set any children", func() {
				Expect(nodes[0].Children).To(BeNil())
				Expect(nodes[1].Children).To(BeNil())
			})

			It("doesn't set any parent", func() {
				Expect(nodes[0].Parent).To(BeNil())
				Expect(nodes[1].Parent).To(BeNil())
			})

		})

	})

})
