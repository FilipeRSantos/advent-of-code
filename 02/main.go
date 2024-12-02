package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/FilipeRSantos/advent-of-code/maths"
)

//go:embed input.txt
var s string

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

var isLevelValid = func(l []int, ignoreError bool) bool {
	removedBadLevel := !ignoreError
	previousValue := 0
	ascending := true

	for i := 0; i < len(l); i++ {
		current := l[i]

		if i == 0 {
			previousValue = current
			ascending = (current < l[i+1])
			continue
		}

		diff := maths.Abs(previousValue - current)
		if diff < 1 || diff > 3 {
			if removedBadLevel {
				return false
			} else {
				removedBadLevel = true
				continue
			}
		}

		if (ascending && previousValue > current) || (!ascending && previousValue < current) {
			if removedBadLevel {
				return false
			} else {
				removedBadLevel = true
				continue
			}
		}

		previousValue = current
	}

	return true
}

func runStep1(input string) int {
	levels := parse(input)

	safeLevels := 0

	for i := 0; i < len(levels); i++ {
		valid := isLevelValid(levels[i], false)

		if valid {
			safeLevels++
		}
	}

	return safeLevels
}

func runStep2(input string) int {
	levels := parse(input)

	safeLevels := 0

	for i := 0; i < len(levels); i++ {
		valid := isLevelValid(levels[i], true)
		if !valid {
			slices.Reverse(levels[i])
			valid = isLevelValid(levels[i], true)
		}

		if valid {
			safeLevels++
		}
	}

	return safeLevels
}

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	output := make([][]int, len(lines))

	for i, line := range lines {
		levels := strings.Split(line, " ")
		output[i] = make([]int, len(levels))
		for j, x := range levels {
			output[i][j] = maths.ParseInt(x)
		}
	}

	return output
}
