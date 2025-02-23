package domain

type PackageResult struct {
	Name     string
	Tests    map[string]*TestResult
	Coverage float64
	Status   string // "pass" or "fail"
	Elapsed  string
}

func (pr *PackageResult) Passed() int {
	passed := 0
	for _, t := range pr.Tests {
		if t.Status == "pass" {
			passed++
		}
	}
	return passed
}

func (pr *PackageResult) Failed() int {
	failed := 0
	for _, t := range pr.Tests {
		if t.Status == "fail" {
			failed++
		}
	}
	return failed
}

func (pr *PackageResult) Skipped() int {
	skipped := 0
	for _, t := range pr.Tests {
		if t.Status == "skip" {
			skipped++
		}
	}
	return skipped
}
