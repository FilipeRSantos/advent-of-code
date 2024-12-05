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

func (c *Config) areUpdatesValid(updateIndex, updatesIndex int) (bool, int) {
	rules, exists := c.rules[c.updates[updatesIndex][updateIndex]]
	if !exists {
		return true, -1
	}

	for _, successor := range rules {
		for j := updateIndex; j >= 0; j-- {
			if c.updates[updatesIndex][j] == successor {
				return false, j
			}
		}
	}

	return true, -1
}

func runStep1(input string) int {
	config := parse(input)

	acc := 0

	for i, updates := range config.updates {
		validRules := true
		for j := range updates {
			if ok, _ := config.areUpdatesValid(j, i); !ok {
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
	config := parse(input)

	acc := 0

	for i, updates := range config.updates {
		adjusted := false

		for {
			invalidRulesPresent := false

			for j := range updates {
				ok, conflictingRule := config.areUpdatesValid(j, i)
				if !ok {
					a := config.updates[i][j]
					b := config.updates[i][conflictingRule]

					config.updates[i][j] = b
					config.updates[i][conflictingRule] = a

					invalidRulesPresent = true
					adjusted = true
				} else {
					continue
				}
			}

			if !invalidRulesPresent {
				break
			}
		}

		if adjusted {
			acc += updates[len(updates)/2]
		}
	}

	return acc
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
