package run

import (
	"fmt"
	"os"
)

func RunModTidy() {
	_, err := run("go mod tidy", true)
	if err != nil {
		fmt.Println("Failed to tidy go modules:", err)
		os.Exit(1)
	}
}
