package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"

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
	return parse(input)
}

func runStep2(input string) int {
	var regex = regexp.MustCompile(`(?s)(don't\(\)).*?(do\(\)|$)`)
	return parse(regex.ReplaceAllString(input, `|`))
}

func parse(input string) int {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	groups := r.FindAllStringSubmatch(input, -1)
	product := 0

	for _, expression := range groups {
		product += maths.ParseInt(expression[1]) * maths.ParseInt(expression[2])
	}

	return product
}
