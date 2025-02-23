package internal

import (
	"log"
	"os"
	"text/template"
)

func WriteOutput() {
	templateContent, err := os.ReadFile("template/output.md")
	if err != nil {
		log.Printf("Failed to read template: %v", err)
		return
	}
	fd := string(templateContent)

	tpl, err := template.New("feed").Parse(fd)
	if err != nil {
		log.Printf("Failed to build template: %v", err)
	}

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
