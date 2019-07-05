package parser

type Node struct {
	Argv     []string
	Children []*Node
	Command  string
	ExitCode int
	Parent   *Node
	Pid      uint64
	Ppid     uint64
}
