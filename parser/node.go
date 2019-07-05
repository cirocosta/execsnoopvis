package parser

type Node struct {
	Pid uint

	Argv     []string
	Command  string
	ExitCode int
	Ppid     uint

	Parent   *Node
	Children []*Node
}
