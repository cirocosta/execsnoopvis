package main

import (
	"context"

	"github.com/cirocosta/execsnoopvis/tracer"
)

type traceCommand struct{}

func (c traceCommand) Execute(args []string) (err error) {
	t := tracer.New()

	err = t.Run(context.TODO())
	return
}
