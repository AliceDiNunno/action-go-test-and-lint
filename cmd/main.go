package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	lintArgs := os.Getenv("INPUT_LINT-ARGS")

	if err := runCommand("golangci-lint", "run", "./...", lintArgs); err != nil {
		log.Printf("golangci-lint found issues or failed: %v", err)
	}

	if err := runCommand("go", "test", "./...", "-v"); err != nil {
		log.Printf("go test failed: %v", err)
	}

	githubOutput := os.Getenv("GITHUB_STEP_SUMMARY")
	if githubOutput != "" {
		f, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Printf("Failed to set output: %v", err)
		} else {
			defer f.Close()
			_, _ = f.WriteString("lint-result=All checks done!\n")
		}
	}

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
