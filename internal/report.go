package internal

import (
	_ "embed"
	"log"
	"os"
	"text/template"

	"github.com/AliceDiNunno/action-go-test-and-lint/domain"
	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

//go:embed templates/output.md
var outputTemplate string

type Report struct {
	packageCoverage map[string]*Coverage
	fileCoverage    map[string]*Coverage
	lint            run.LintResult
	testsResult     map[string]*domain.PackageResult
}

func WriteReport(report Report) {
	tpl := template.Must(template.New("feed").Parse(outputTemplate))

	githubOutput := os.Getenv("GITHUB_STEP_SUMMARY")
	if githubOutput != "" {
		f, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Printf("Failed to set output: %v", err)
		} else {
			defer f.Close()
			tpl.Execute(f, report)
		}
	}
}
