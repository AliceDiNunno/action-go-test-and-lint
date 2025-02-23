package internal

import (
	_ "embed"
	"log"
	"os"
	"text/template"
)

//go:embed templates/output.md
var outputTemplate string

func WriteOutput() {
	tpl := template.Must(template.New("feed").Parse(outputTemplate))

	githubOutput := os.Getenv("GITHUB_STEP_SUMMARY")
	if githubOutput != "" {
		f, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Printf("Failed to set output: %v", err)
		} else {
			defer f.Close()
			tpl.Execute(f, nil)
		}
	}
}
