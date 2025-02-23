package internal

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

func RunUnit() {
	data, err := run("go test ./... -json -coverprofile=coverage.out", true)
	if err != nil {
		log.Printf("go test failed: %v", err)
	}

	spew.Dump(data)
}
