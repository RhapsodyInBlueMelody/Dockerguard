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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dockguard <Dockerfile-path>")
		return
	}
	var path string = os.Args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("%s[ERROR] Can't read %q: %v%s\n",
			Red, path, err, Reset)
		return
	}
	fmt.Println("Scanning...", path)

	lines := strings.Split(string(data), "\n")

	for index, line := range lines {
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		if !isFromInstruction(parts[0]) {
			continue
		}

		fmt.Println("[LOG] Found a FROM instruction at line ", index+1, ":", line)
		if len(parts) < 2 {
			fmt.Printf("%s[WARNING] Line %d: FROM instruction is missing an image name%s\n",
				Yellow, index+1, Reset)
			continue
		}

		if checkImageTag(line) == 1 {
			fmt.Printf("%s[WARNING] Line %d: FROM instruction is missing an image tag%s\n",
				Yellow, index+1, Reset)
			continue
		} else if checkImageTag(line) == 2 {
			fmt.Printf("%s[WARNING] Line %d: FROM instruction using latest tag%s\n",
				Yellow, index+1, Reset)
			continue
		}
	}
}
