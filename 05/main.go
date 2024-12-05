package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/FilipeRSantos/advent-of-code/maths"
)

//go:embed input.txt
var s string

type Config struct {
	rules   map[int][]int
	updates [][]int
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

func (c *Config) areUpdatesValid(updateIndex, page, updatesIndex int) bool {
	rules, exists := c.rules[page]
	if !exists {
		return true
	}

	for _, successor := range rules {
		for j := updateIndex; j >= 0; j-- {
			if c.updates[updatesIndex][j] == successor {
				return false
			}
		}
	}

	return true
}

func runStep1(input string) int {
	config := parse(input)

	acc := 0

	for i, updates := range config.updates {
		validRules := true
		for j, page := range updates {
			if !config.areUpdatesValid(j, page, i) {
				validRules = false
				break
			}
		}

		if validRules {
			acc += updates[len(updates)/2]
		}
	}

	return acc
}

func runStep2(input string) int {
	return 0
}

func parse(input string) *Config {
	rulesSection := true
	lines := strings.Split(input, "\n")

	rules := make(map[int][]int, 10)
	updatesIndex := 0
	var updates [][]int

	for i, line := range lines {
		if line == `` {
			rulesSection = false
			updates = make([][]int, len(lines)-i-1)
			continue
		}

		if rulesSection {
			parsedRules := strings.Split(line, "|")

			previous := maths.ParseInt(parsedRules[0])
			successor := maths.ParseInt(parsedRules[1])

			currentRules, exists := rules[previous]
			if !exists {
				rules[previous] = []int{successor}
			} else {
				currentRules = append(currentRules, successor)
				rules[previous] = currentRules
			}
		} else {
			currentUpdate := strings.Split(line, ",")
			updates[updatesIndex] = make([]int, len(currentUpdate))
			for i, x := range currentUpdate {
				updates[updatesIndex][i] = maths.ParseInt(x)
			}
			updatesIndex++
		}
	}

	return &Config{
		rules,
		updates,
	}
}
