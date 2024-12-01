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
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += maths.Abs(left[i] - right[i])
	}
	return sum
}

func runStep2(input string) int {
	left, right := parse(input)
	frequency := 0
	current := 0
	similarity := 0
	lastRightIndex := 0

	for i := 0; i < len(left); i++ {

		if current == left[i] {
			similarity += current * frequency
			continue
		} else {
			current = left[i]
			frequency = 0
		}

		for j := lastRightIndex; j < len(right); j++ {
			if right[j] == current {
				frequency++
			} else if right[j] > current {
				lastRightIndex = j
				break
			}
		}

		similarity += current * frequency
		current = left[i]
	}

	return similarity
}

func parse(input string) ([]int, []int) {
	var lColumn, rColumn []int

	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, " ")
		lColumn = append(lColumn, maths.ParseInt(values[0]))
		rColumn = append(rColumn, maths.ParseInt(values[len(values)-1]))
	}

	slices.Sort(lColumn)
	slices.Sort(rColumn)

	return lColumn, rColumn
}
