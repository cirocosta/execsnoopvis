package parser

import (
	"fmt"
)

type Node struct {
	Argv     []string
	Children []*Node
	Command  string
	ExitCode int
	Parent   *Node
	Pid      uint64
	Ppid     uint64
}

func (n *Node) Name() string {
	return fmt.Sprintf("%s [%d]", n.Command, n.Pid)
}
