package internal

import (
	"fmt"

	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

func RunTests() (Report, bool) {
	unitTests, unitSuccess := run.RunUnit()
	lint, lintSuccess := run.RunLint()
	packageCoverage, fileCoverage, totalCoverage := RunCoverage()

	fmt.Printf("Unit tests: %v\n", unitSuccess)
	fmt.Printf("Lint: %v\n", lintSuccess)

	return Report{
		TotalCoverage:   totalCoverage,
		PackageCoverage: packageCoverage,
		FileCoverage:    fileCoverage,
		Lint:            lint,
		TestsResult:     MapTestRawOutputToPackageResults(unitTests),
	}, unitSuccess && lintSuccess
}
