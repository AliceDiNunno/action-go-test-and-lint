package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Coverage struct {
	Statements int
	Covered    int
}

func RunCoverage() (coverageByPackage map[string]*Coverage, coverageByFile map[string]*Coverage, totalCoverage *Coverage) {
	file, err := os.Open("coverage.out")
	if err != nil {
		fmt.Println("Error opening coverage file:", err)
		return nil, nil, nil
	}
	defer file.Close()

	coverageByFile = make(map[string]*Coverage)
	coverageByPackage = make(map[string]*Coverage)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mode:") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		regex := regexp.MustCompile(`([\w./-]+)/[\w-]+\.go`)
		fullPathRegex := regexp.MustCompile(`([\w./-]+\.go)`)

		// Find the package name using the regex
		match := regex.FindStringSubmatch(fields[0])
		packageName := ""
		if len(match) > 1 {
			packageName = match[1]
		}

		// Find the package name using the regex
		match = fullPathRegex.FindStringSubmatch(fields[0])
		fileName := ""
		if len(match) > 1 {
			fileName = match[1]
		}

		statements := 1
		covered := 0
		if fields[2] == "1" {
			covered = 1
		}

		if _, exists := coverageByFile[fileName]; !exists {
			coverageByFile[fileName] = &Coverage{}
		}

		if _, exists := coverageByPackage[packageName]; !exists {
			coverageByPackage[packageName] = &Coverage{}
		}

		coverageByPackage[packageName].Statements += statements
		coverageByPackage[packageName].Covered += covered

		coverageByFile[fileName].Statements += statements
		coverageByFile[fileName].Covered += covered
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading coverage file:", err)
		return nil, nil, nil
	}

	totalStatements := 0
	totalCovered := 0

	for _, coverage := range coverageByFile {
		totalStatements += coverage.Statements
		totalCovered += coverage.Covered
	}

	return coverageByPackage, coverageByFile, &Coverage{
		Statements: totalStatements,
		Covered:    totalCovered,
	}
}
