package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/FilipeRSantos/advent-of-code/maths"
)

//go:embed input.txt
var s string
var values map[int][]int

func main() {
	var ans int
	args := os.Args[1]

	if args == "1" {
		ans = runStep1(s, 25)
	} else {
		ans = runStep1(s, 75)
	}

	fmt.Println("Output: ", ans)
}

func applyStoneRules(n int) []int {
	if n == 0 {
		return []int{1}
	}

	s := strconv.Itoa(n)
	if len(s)%2 == 0 {
		leftSide := maths.ParseInt(s[:len(s)/2])
		rightSide := maths.ParseInt(s[len(s)/2:])

		return []int{leftSide, rightSide}
	}

	return []int{n * 2024}

}

func computeNumber(n int) []int {
	value, exists := values[n]
	if !exists {
		value = applyStoneRules(n)
		values[n] = value
	}
	return value
}

func computeNSteps(n, steps int) int {

	if steps <= 0 {
		return 0
	}

	values := computeNumber(n)
	newItems := len(values) - 1
	for _, curr := range values {
		newItems += computeNSteps(curr, steps-1)
	}

	return newItems
}

func runStep1(input string, blinks int) int {
	stones := parse(input)
	values = make(map[int][]int, math.MaxUint16)

	acc := len(stones)
	for _, stone := range stones {
		acc += computeNSteps(stone, blinks)
	}
	return acc
}

func parse(input string) []int {
	stones := strings.Split(input, " ")
	arr := make([]int, len(stones))

	for i, stone := range stones {
		arr[i] = maths.ParseInt(string(stone))
	}

	return arr
}
