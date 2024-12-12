package main

import (
	"fmt"
	"testing"
)

func TestStep1SampleInput(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{`AAAA
BBCD
BBCC
EEEC`, 140},
		{`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`, 772},
		{`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`, 1930},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runStep1(tt.input)
			if ans != tt.expected {
				t.Errorf(`TestStep1SampleInput got %d, expected %d`, ans, tt.expected)
			}
		})
	}
}
