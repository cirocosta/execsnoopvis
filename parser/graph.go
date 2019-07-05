package parser

// PopulateNodes mutates the nodes that are passed in as arguments so that they
// contain information about other nodes in the array that form relationships.
//
func PopulateNodes(nodes []*Node) {
	perPidNodeMapping := make(map[uint64]*Node, len(nodes))

	for _, node := range nodes {
		perPidNodeMapping[node.Pid] = node
	}

	for _, node := range perPidNodeMapping {
		node.Parent = perPidNodeMapping[node.Ppid]

		if node.Parent != nil {
			node.Parent.Children = append(
				node.Parent.Children,
				node,
			)
		}
	}

	return
}

// FindRoots searches for those nodes in a set of nodes who are roots of trees.
//
func FindRoots(nodes []*Node) (roots []*Node) {
	for _, node := range nodes {
		if node.Parent != nil {
			continue
		}

		roots = append(roots, node)
	}

	return
}
