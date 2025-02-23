package main

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/AliceDiNunno/action-go-test-and-lint/internal"
)

func main() {
	spew.Dump(internal.RunUnit())
	spew.Dump(internal.RunLint())
	internal.WriteOutput()
}
