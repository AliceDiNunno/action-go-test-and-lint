package main

import (
	"github.com/AliceDiNunno/action-go-test-and-lint/internal"
	"github.com/AliceDiNunno/action-go-test-and-lint/internal/run"
)

func main() {
	//dump all env variables

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	run.RunModTidy()

	report, success := internal.RunTests()
	internal.WriteReport(report)

	if !success {
		println("Tests failed")
	}
}
