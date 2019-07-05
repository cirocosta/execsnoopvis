package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	input = flag.String("input", "", "location of the file containing the output of execsnoop")
)

func main() {
	flag.Parse()

	if *input == "" {
		fmt.Println("the `input` flag must be non-empty.")
		os.Exit(1)
		return
	}
}
