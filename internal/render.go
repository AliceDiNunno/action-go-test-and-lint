package internal

import (
	"log"
	"os"
)

func WriteOutput() {
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
}
