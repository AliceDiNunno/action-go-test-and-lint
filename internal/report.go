package internal

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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

type PieData struct {
	Field string
	Value int
	Color string
}

func buildPie(data []PieData) string {
	//mermaid pie
	pie := "pie\n"

	//sort pie by value
	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})

	//remove all values that are 0
	for i := 0; i < len(data); i++ {
		if data[i].Value == 0 {
			data = append(data[:i], data[i+1:]...)
			i--
		}
	}

}

func testPie(pkg *domain.PackageResult) string {
	return ""
}

func coveragePie(coverage *Coverage) string {
	return ""
}

func WriteReport(report Report) {
	repository := os.Getenv("GITHUB_REPOSITORY")
	blob := os.Getenv("GITHUB_SHA")

	funcMap := template.FuncMap{
		"substract": func(a float64, b float64) float64 { return a - b },
		"percent":   func(a int, b int) float64 { return float64(a) / float64(b) * 100 },
		"trim":      func(f float64) string { return fmt.Sprintf("%.2f", f) },
		"link": func(fileName string, line int) string {
			return fmt.Sprintf("https://github.com/%s/blob/%s/%s#L%d", repository, blob, fileName, line)
		},
		"pkgFailedCount":  func(pkg *domain.PackageResult) int { return pkg.Failed() },
		"pkgSkippedCount": func(pkg *domain.PackageResult) int { return pkg.Skipped() },
		"pkgPassedCount":  func(pkg *domain.PackageResult) int { return pkg.Passed() },
		"pkgDuration":     func(pkg *domain.PackageResult) string { return pkg.Elapsed },
		"testOutput":      func(test *domain.TestResult) string { return strings.Trim(test.Output, "\n") },
		"detailOpened": func(str string) string {
			if len(strings.Split(strings.Trim(str, "\n"), "\n")) > 2 {
				return "open"
			}
			return ""
		},
		"pkgBadge":  func(pkg *domain.PackageResult) string { return pkg.Badge() },
		"testBadge": func(test *domain.TestResult) string { return test.Badge() },
		"totalPassed": func() int {
			passed := 0
			for _, pkg := range report.TestsResult {
				passed += pkg.Passed()
			}
			return passed
		},
		"totalFailed": func() int {
			failed := 0
			for _, pkg := range report.TestsResult {
				failed += pkg.Failed()
			}
			return failed
		},
		"totalSkipped": func() int {
			skipped := 0
			for _, pkg := range report.TestsResult {
				skipped += pkg.Skipped()
			}
			return skipped
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
