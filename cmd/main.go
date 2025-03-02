package main

import (
	"fmt"
	"os"

	"github.com/AliceDiNunno/action-go-test-and-lint/internal"
)

func main() {
	//dump all env variables
	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	report, success := internal.RunTests()
	internal.WriteReport(report)

	if !success {
		println("Tests failed")
		os.Exit(1)
	}
}
