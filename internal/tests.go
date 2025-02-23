package internal

import (
	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

func RunTests() (Report, bool) {
	unitTests, unitSuccess := run.RunUnit()
	lint, lintSuccess := run.RunLint()
	packageCoverage, fileCoverage := RunCoverage()

	return Report{
		packageCoverage: packageCoverage,
		fileCoverage:    fileCoverage,
		lint:            lint,
		testsResult:     MapTestRawOutputToPackageResults(unitTests),
	}, unitSuccess && lintSuccess
}
