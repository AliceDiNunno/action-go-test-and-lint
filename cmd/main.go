package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/davecgh/go-spew/spew"

	"github.com/AliceDiNunno/action-go-test-and-lint/internal"
)

func main() {
	lintArgs := os.Getenv("INPUT_LINT-ARGS")

	if err := runCommand("golangci-lint", "run", "./...", lintArgs); err != nil {
		log.Printf("golangci-lint found issues or failed: %v", err)
	}

	spew.Dump("hello")

	internal.RunUnit()
	internal.WriteOutput()

	fmt.Println("Done. Exiting.")
}

func runCommand(name string, args ...string) error {
	var finalArgs []string
	for _, a := range args {
		if a != "" {
			finalArgs = append(finalArgs, a)
		}
	}
	cmd := exec.Command(name, finalArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Running command: %s %v", name, finalArgs)
	return cmd.Run()
}
