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

func isFromInstruction(param1 string) bool {
	parts := strings.Fields(param1)
	if len(parts) == 0 {
		return false
	}
	return strings.EqualFold(parts[0], "FROM")
}

func useLatestTag(param1 string) bool {
	parts := strings.Split(param1, ":")
	if len(parts) <= 1 {
		return false
	}
	return parts[1] == "latest"
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
		if useLatestTag(parts[1]) {
			fmt.Println("[WARNING] Line ", index+1, ": base image uses the explicit latest tag: ", line)
		}
	}
}
