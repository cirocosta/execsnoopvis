package parser

import (
	"github.com/awalterschulze/gographviz"
	"github.com/pkg/errors"
)

func must(err error) {
	if err == nil {
		return
	}

	panic(err)
}

const GRAPH_NAME = "G"

// Render renders the trees described by the `roots` argument to `writer` using
// the `dot` syntax.
//
func Render(roots []*Node) (contents string, err error) {
	if len(roots) == 0 {
		err = errors.Errorf("number of roots must be >= 1")
		return
	}

	graph := gographviz.NewGraph()

	must(graph.SetName(GRAPH_NAME))
	must(graph.SetDir(true))

	for _, node := range roots {
		fillGraph(node, graph)
	}

	contents = graph.String()
	return
}

// fillFraph fills a given graph with the nodes and edges that pertain to that
// tree.
//
func fillGraph(root *Node, graph *gographviz.Graph) {
	src := root.Name()

	must(graph.AddNode(GRAPH_NAME, src, nil))

	for _, child := range root.Children {
		dst := child.Name()

		must(graph.AddEdge(src, dst, true, nil))

		fillGraph(child, graph)
	}

	return
}
