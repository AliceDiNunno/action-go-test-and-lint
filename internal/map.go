package internal

import (
	"fmt"
	"strings"
	"time"

	"github.com/AliceDiNunno/action-go-test-and-lint/domain"
	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

func MapTestRawOutputToPackageResults(raw []run.TestRawOutput) map[string]*domain.PackageResult {
	packages := make(map[string]*domain.PackageResult)

	for _, line := range raw {
		testName := line.Test

		if _, ok := packages[line.Package]; !ok {
			packages[line.Package] = &domain.PackageResult{
				Name:   line.Package,
				Tests:  make(map[string]*domain.TestResult),
				Status: "pass",
			}
		}
		pkg := packages[line.Package]

		if testName != "" {
			if _, ok := pkg.Tests[testName]; !ok {
				pkg.Tests[testName] = &domain.TestResult{
					Name:   testName,
					Status: "",
				}
			}

			if line.Output != "" {
				// Append output to the test
				pkg.Tests[testName].Output = pkg.Tests[testName].Output + line.Output
			}
		}

		if line.Elapsed != nil {
			elapsed := time.Duration(*line.Elapsed*1000) * time.Millisecond
			pkg.Elapsed = fmt.Sprintf("%v", elapsed)
		}

		switch strings.ToLower(line.Action) {
		case "pass":
			if testName == "" {
				if pkg.Status != "fail" {
					pkg.Status = "pass"
				}
			} else {
				tr := pkg.Tests[testName]
				if tr.Status == "" || tr.Status == "pass" {
					tr.Status = "pass"
				}
			}

		case "fail":
			if testName == "" {
				pkg.Status = "fail"
			} else {
				tr := pkg.Tests[testName]
				tr.Status = "fail"
				pkg.Status = "fail"
			}

		case "skip":
			if testName == "" {
				pkg.Status = "fail"
			} else {
				tr := pkg.Tests[testName]
				if tr.Status == "" {
					tr.Status = "skip"
				}
			}
		}
	}

	for _, pkg := range packages {
		if len(pkg.Tests) == 0 {
			pkg.Status = "skip"
		}
	}

	// If any test is still "" status after all lines, let's default it to "pass" (or "skip"):
	for _, pkg := range packages {
		for _, t := range pkg.Tests {
			if t.Status == "" {
				t.Status = "pass"
			}
		}
	}

	return packages
}
