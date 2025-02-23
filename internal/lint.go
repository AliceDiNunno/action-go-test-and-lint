package internal

import (
	"encoding/json"
	"log"
)

type LintResult struct {
	Issues []struct {
		FromLinter string
		Text       string
		Pos        struct {
			Filename string
			Line     int
			Column   int
		}
	}
	Report struct {
		Warnings []struct {
			Tag  string
			Text string
		}
	}
}

func RunLint() (results LintResult, success bool) {
	data, err := run("golangci-lint run ./... --out-format json", false)
	success = err == nil

	if err != nil {
		log.Printf("golangci-lint failed: %v", err)
	}

	err = json.Unmarshal([]byte(data), &results)
	if err != nil {
		log.Printf("Failed to parse golangci-lint output: %v", err)
		success = false
	}

	return
}
