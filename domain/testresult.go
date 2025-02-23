package domain

type TestResult struct {
	Name   string
	Status string
	Output string
}

func (pr *TestResult) Badge() string {
	switch pr.Status {
	case "pass":
		return "🟢"
	case "fail":
		return "🔴"
	case "skip":
		return "🟡"
	}
	return ""
}
