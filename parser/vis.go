package parser

import (
	"io"
	// "strconv"

	"github.com/awalterschulze/gographviz"
	// "github.com/pkg/errors"
)

func must(err error) {
	if err == nil {
		return
	}

	panic(err)
}

const GRAPH_NAME = "graph"

// Render renders the trees described by the `roots` argument to `writer` using
// the `dot` syntax.
//
func Render(roots []*Node, writer io.Writer) (err error) {
	// if len(roots) == 0 {
	// 	err = errors.Errorf("number of roots must be >= 1")
	// 	return
	// }

	// graph := gographviz.NewGraph()

	// must(graph.SetName(GRAPH_NAME))
	// must(graph.SetDir(true))

	// firstRoot := roots[0]

	return
}

// FillFraph fills a given graph with the nodes and edges that pertain to that
// tree.
//
func FillGraph(root *Node, graph *gographviz.Graph) {
	// src := strconv.Itoa(root.Pid)

	// must(graph.AddNode(GRAPH_NAME, src))

	// for _, child := range root.Children {
	// 	dst := strconv.Itoa(child.Pid)

	// 	must(graph.AddEdge(src, dst, true, nil))

	// 	FillGraph(child, graph)
	// }

	return
}
