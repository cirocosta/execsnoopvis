package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cirocosta/execsnoopvis/parser"
)

var (
	input = flag.String("input", "", "location of the file containing the output of execsnoop")
)

func mainWithExitCode() (statusCode int) {
	flag.Parse()

	if *input == "" {
		fmt.Println("the `input` flag must be non-empty.")
		statusCode = 1
		return
	}

	file, err := os.Open(*input)
	if err != nil {
		fmt.Println("failed opening file %s - %+v\n", *input, err)
		statusCode = 1
		return
	}

	defer file.Close()

	scanner := parser.NewScanner(file)

	nodes, err := scanner.ScanAll()
	if err != nil {
		fmt.Println("failed while scanning file: %v", err)
		statusCode = 1
		return
	}

	parser.PopulateNodes(nodes)
	roots := parser.FindRoots(nodes)

	err = parser.Render(roots, os.Stdout)
	if err != nil {
		fmt.Println("failed rendering graph: %v", err)
		statusCode = 1
		return
	}

	return
}

func main() {
	os.Exit(mainWithExitCode())
}
