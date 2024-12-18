package maths

import (
	"math"
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

func Max(n, m int) int {
	if n < m {
		return m
	}

	return n
}

func Min(n, m int) int {
	if n > m {
		return m
	}

	return n
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
