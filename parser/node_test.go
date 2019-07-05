package parser_test

import (
	"github.com/cirocosta/execsnoopvis/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Node", func() {

	type Case struct {
		Node     *parser.Node
		Expected string
	}

	DescribeTable("Name()",
		func(c Case) {
			Expect(c.Expected).To(Equal(c.Node.Name()))
		},
		Entry("nothing set", Case{
			Node:     &parser.Node{},
			Expected: " [0]",
		}),
		Entry("without command", Case{
			Node: &parser.Node{
				Pid: 123,
			},
			Expected: " [123]",
		}),
		Entry("without pid", Case{
			Node: &parser.Node{
				Command: "go",
			},
			Expected: "go [0]",
		}),
		Entry("with all set", Case{
			Node: &parser.Node{
				Command: "go",
				Pid:     123,
			},
			Expected: "go [123]",
		}),
	)

})
