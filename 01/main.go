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

func runStep1(input string) int {
	left, right := parse(input)

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += maths.Abs(left[i] - right[i])
	}
	return sum
}

func runStep2(input string) int {
	left, right := parse(input)
	similarity := 0
	rightNumbers := map[int]int{}

	for i := 0; i < len(right); i++ {
		v, exists := rightNumbers[right[i]]
		if !exists {
			rightNumbers[right[i]] = 1
		} else {
			rightNumbers[right[i]] = v + 1
		}
	}

	for i := 0; i < len(left); i++ {
		if v, exists := rightNumbers[left[i]]; exists {
			similarity += left[i] * v
		}
	}

	return similarity
}

func parse(input string) ([]int, []int) {
	var lColumn, rColumn []int

	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, "   ")
		lColumn = append(lColumn, maths.ParseInt(values[0]))
		rColumn = append(rColumn, maths.ParseInt(values[1]))
	}

	return lColumn, rColumn
}
