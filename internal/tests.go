package internal

import (
	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

func RunTests() (Report, bool) {
	unitTests, unitSuccess := run.RunUnit()
	lint, lintSuccess := run.RunLint()
	packageCoverage, fileCoverage := RunCoverage()

	return Report{
		PackageCoverage: packageCoverage,
		FileCoverage:    fileCoverage,
		Lint:            lint,
		TestsResult:     MapTestRawOutputToPackageResults(unitTests),
	}, unitSuccess && lintSuccess
}
