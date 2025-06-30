package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	envFile := ".env"
	if len(os.Args) > 1 {
		envFile = os.Args[1]
	}

	file, err := os.Open(envFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			value = strings.Trim(value, `"'`)
			fmt.Printf("export %s='%s'\n", key, strings.ReplaceAll(value, "'", "'\\''"))
		}
	}
}
