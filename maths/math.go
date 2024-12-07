package maths

import (
	"strconv"
)

func ParseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}
