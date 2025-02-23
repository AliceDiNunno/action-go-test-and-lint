package internal

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/AliceDiNunno/action-go-test-and-lint/domain"
	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

//go:embed templates/output.md
var outputTemplate string

type Report struct {
	PackageCoverage map[string]*Coverage
	FileCoverage    map[string]*Coverage
	Lint            run.LintResult
	TestsResult     map[string]*domain.PackageResult
	TotalCoverage   *Coverage
}

// https://github.com/yeencloud/test/blob/82a6c99cdf8ff17de2fa466ce48ec9fb2da38f2f/strings.go#L8
func WriteReport(report Report) {
	repository := os.Getenv("GITHUB_REPOSITORY")
	blob := os.Getenv("GITHUB_SHA")

	funcMap := template.FuncMap{
		"substract": func(a float64, b float64) float64 { return a - b },
		"percent":   func(a int, b int) float64 { return float64(a) / float64(b) * 100 },
		"trim":      func(f float64) string { return fmt.Sprintf("%.2f", f) },
		"link": func(fileName string, line int) string {
			return fmt.Sprintf("https://github.com/%sblob%s/%s#L%d", repository, blob, fileName, line)
		},
	}
	tpl := template.Must(template.New("feed").Funcs(funcMap).Parse(outputTemplate))

	githubOutput := os.Getenv("GITHUB_STEP_SUMMARY")
	if githubOutput != "" {
		f, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Printf("Failed to set output: %v", err)
		} else {
			defer f.Close()
			err := tpl.Execute(f, report)
			if err != nil {
				log.Printf("Failed to write report: %v", err)
				return
			}
		}
	}
}
