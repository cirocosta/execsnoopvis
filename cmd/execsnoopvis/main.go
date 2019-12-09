package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

var execsnoopvis struct {
	Trace traceCommand `command:"trace"`
}

func main() {
	parser := flags.NewParser(&execsnoopvis, flags.HelpFlag|flags.PassDoubleDash)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
