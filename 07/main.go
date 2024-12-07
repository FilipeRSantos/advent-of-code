package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FilipeRSantos/advent-of-code/maths"
	"github.com/FilipeRSantos/advent-of-code/utils"
)

//go:embed input.txt
var s string

type Equation struct {
	total   int
	numbers []int
}

func (e *Equation) isViable(part1 bool) bool {
	var permutationBase int
	if part1 {
		permutationBase = 2
	} else {
		permutationBase = 3
	}
	maxPermutation := utils.LeftPadWith(strconv.Itoa(permutationBase-1), strconv.Itoa(permutationBase-1), len(e.numbers)-1)

	result, _ := strconv.ParseInt(maxPermutation, permutationBase, 64)

	for i := int64(0); i <= result; i++ {
		acc := 0
		permutation := utils.LeftPadWith(strconv.FormatInt(int64(i), permutationBase), "0", len(maxPermutation))

		if len(permutation) != len(e.numbers)-1 {
			log.Fatalf("permutation %s does not follow the expected pattern for %d ", permutation, e.total)
		}

		for i, char := range permutation {
			if i == 0 {
				acc = e.numbers[i]
			}

			if char == '0' {
				acc += e.numbers[i+1]
			} else if char == '1' {
				acc *= e.numbers[i+1]
			} else {
				acc = maths.ParseInt(strconv.Itoa(acc) + strconv.Itoa(e.numbers[i+1]))
			}
		}

		if acc == e.total {
			return true
		}
	}
	return false
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

func parse(input string) []Equation {
	lines := strings.Split(input, "\n")
	output := make([]Equation, len(lines))

	for i, line := range lines {
		section := strings.Split(line, ":")
		total := maths.ParseInt(section[0])
		values := strings.Split(strings.Trim(section[1], " "), " ")
		numbers := make([]int, len(values))
		for j, value := range values {
			numbers[j] = maths.ParseInt(value)
		}

		output[i] = Equation{
			total,
			numbers,
		}

	}

	return output
}

func runStep1(input string) int {
	equations := parse(input)
	calibration := 0
	for _, equation := range equations {
		if equation.isViable(true) {
			calibration += equation.total
		}
	}

	return calibration
}

func runStep2(input string) int {
	equations := parse(input)
	calibration := 0
	for _, equation := range equations {
		if equation.isViable(false) {
			calibration += equation.total
		}
	}

	return calibration
}
