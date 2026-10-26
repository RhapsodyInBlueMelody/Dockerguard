package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

const (
	TagMissing int = 1
	LatestTag  int = 2
	VersionTag int = 3
)

type Warning struct {
	Line    int
	Type    int
	Message string
}

type Log struct {
	Line    int
	Message string
}

func isFromInstruction(param1 string) bool {
	parts := strings.Fields(param1)
	if len(parts) == 0 {
		return false
	}
	return strings.EqualFold(parts[0], "FROM")
}

func checkImageTag(param1 string) int {
	colonIndex := strings.LastIndex(param1, ":")
	slashIndex := strings.LastIndex(param1, "/")
	if colonIndex < 0 || colonIndex < slashIndex {
		return TagMissing
	}
	if param1[colonIndex+1:] == "latest" {
		return LatestTag
	}
	return VersionTag

}

func fromImageParser(parts []string, index int) (Warning, bool) {

	if len(parts) < 2 {
		warning := Warning{
			Line:    index + 1,
			Type:    TagMissing,
			Message: "FROM instruction is missing an image tag",
		}

		return warning, true

	}

	var status int = checkImageTag(parts[1])

	switch status {
	case TagMissing:
		warning := Warning{
			Line:    index + 1,
			Type:    TagMissing,
			Message: "FROM instruction is missing an image tag",
		}
		return warning, true

	case LatestTag:
		warning := Warning{
			Line:    index + 1,
			Type:    LatestTag,
			Message: "FROM instruction using latest tag",
		}
		return warning, true
	}

	return Warning{}, false
}

func parserDockerfile(data []byte) int {

	lines := strings.Split(string(data), "\n")
	var warningCount int = 0
	for index, line := range lines {
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		result, foundWarning := fromImageParser(parts, index)

		if foundWarning {
			fmt.Printf("[WARNING] Line %d: %s\n", result.Line, result.Message)
			warningCount++
		}

	}

	return warningCount
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: dockguard <Dockerfile-path>")
		os.Exit(2)
	}
	var path string = os.Args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("%s[ERROR] Can't read %q: %v%s\n",
			Red, path, err, Reset)
		os.Exit(2)
	}
	fmt.Println("Scanning...", path)

	var warnings int = parserDockerfile(data)

	fmt.Println("Scanning Finished!")
	fmt.Printf("Warning found: %d\n", warnings)
}
