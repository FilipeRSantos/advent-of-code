package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var s string

var available map[string]bool
var patterns map[string]struct{}
var towels []string

type Coordinate struct {
	x, y int
}

func main() {
	var ans int
	args := os.Args[1]

	if args == "1" {
		ans = runStep1(s)
	} else {
		ans = runStep2(s)
	}

	fmt.Println("Output: ", ans)
}

func isPatternPossible(towel string) bool {
	possible, calculated := available[towel]
	if calculated {
		return possible
	}

	for k := range patterns {
		if !strings.Contains(towel, k) {
			continue
		}

		towelLeft := strings.Replace(towel, k, "-", 1)
		parts := strings.Split(towelLeft, "-")
		possible := isPatternPossible(parts[0])
		if len(parts) > 1 {
			possible = possible && isPatternPossible(parts[1])
		}

		if possible {
			available[towel] = true
			return true
		}
	}

	available[towel] = false
	return false
}

// n < 400
func runStep1(input string) int {
	parse(input)

	acc := 0
	for _, towel := range towels {
		if isPatternPossible(towel) {
			acc++
		}
	}

	return acc
}

func runStep2(input string) int {
	return 0
}

func parse(input string) {
	lines := strings.Split(input, "\n")

	available = make(map[string]bool)
	patterns = make(map[string]struct{})
	towels = make([]string, len(lines)-2)

	available[""] = true
	for _, pattern := range strings.Split(lines[0], ", ") {
		patterns[pattern] = struct{}{}
		available[pattern] = true
	}

	copy(towels, lines[2:])
}
