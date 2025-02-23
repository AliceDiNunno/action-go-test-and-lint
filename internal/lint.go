package internal

import (
	"encoding/json"
	"log"
)

type LintResult struct {
	Issues []struct {
		FromLinter string `json:"from_linter"`
		Text       string `json:"text"`
		Pos        struct {
			Filename string `json:"filename"`
			Line     int    `json:"line"`
			Column   int    `json:"column"`
		} `json:"pos"`
	} `json:"issues"`
	Report struct {
		Warnings []struct {
			Tag  string `json:"tag"`
			Text string `json:"text"`
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
