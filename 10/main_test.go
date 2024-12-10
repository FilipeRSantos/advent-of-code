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
		{`0123
1234
8765
9876`, 1},
		{`..90..9
...1.98
...2..7
6543456
765.987
876....
987....`, 4},
		{`10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`, 3},
		{`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, 36},
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
		{`012345
123456
234567
345678
4.6789
56789.`, 227},
		{`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, 81},
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
