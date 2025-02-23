package domain

type TestResult struct {
	Name   string
	Status string // "pass", "fail", or "skip"
	Output string
}
