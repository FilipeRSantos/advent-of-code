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

func TestStep2SampleInput(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{`AAAA
BBCD
BBCC
EEEC`, 80},
		{`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`, 436},
		{`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`, 236},
		{`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`, 368},
		{`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`, 1206},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runStep2(tt.input)
			if ans != tt.expected {
				t.Errorf(`TestStep2SampleInput got %d, expected %d`, ans, tt.expected)
			}
		})
	}
}
