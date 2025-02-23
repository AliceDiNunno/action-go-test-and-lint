package main

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/AliceDiNunno/action-go-test-and-lint/internal"
	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

func main() {
	run.RunModTidy()
	spew.Dump(run.RunUnit())
	spew.Dump(run.RunLint())
	internal.WriteOutput()
}
