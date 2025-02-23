package main

import (
	"github.com/AliceDiNunno/action-go-test-and-lint/internal"
)

func main() {
	internal.RunUnit()
	internal.RunLint()
	internal.WriteOutput()
}
