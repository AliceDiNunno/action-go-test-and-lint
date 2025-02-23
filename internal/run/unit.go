package run

import (
	"encoding/json"
	"log"
	"strings"
)

type TestRawOutput struct {
	Package string
	Time    string
	Action  string
	Test    string
	Output  string

	Elapsed *float64
}

func RunUnit() (results []TestRawOutput, success bool) {
	data, err := run("go test ./... -json -coverprofile=coverage.out", true)
	if err != nil {
		log.Printf("go test failed: %v", err)
	}

	results = []TestRawOutput{}
	lines := strings.NewReader(data)
	decoder := json.NewDecoder(lines)
	for {
		var result TestRawOutput
		if err := decoder.Decode(&result); err != nil {
			break
		}
		results = append(results, result)
	}
	return
}
