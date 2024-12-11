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
var stones map[int]int

func main() {
	var ans int
	args := os.Args[1]

	if args == "1" {
		ans = runStep(s, 25)
	} else {
		ans = runStep(s, 75)
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

func insertOrAdd(dic map[int]int, n, qtd int) {
	value, exists := dic[n]
	if !exists {
		dic[n] = qtd
	} else {
		dic[n] = value + qtd
	}
}

func runStep(input string, blinks int) int {
	previousStone := make(map[int]int, math.MaxUint16)
	for _, stone := range parse(input) {
		insertOrAdd(previousStone, stone, 1)
	}

	for range blinks {
		stones = make(map[int]int, math.MaxUint16)
		for key, value := range previousStone {
			t := applyStoneRules(key)
			for _, x := range t {
				insertOrAdd(stones, x, value)
			}

		}

		previousStone = make(map[int]int, math.MaxUint16)
		for key, value := range stones {
			previousStone[key] = value
		}

	}

	acc := 0
	for _, value := range stones {
		acc += value
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
