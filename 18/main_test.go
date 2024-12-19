package main

import (
	"fmt"
	"testing"
)

func TestStep1SampleInput(t *testing.T) {
	var tests = []struct {
		input     string
		size      int
		corrupted int
		expected  int
	}{
		{`5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`, 7, 12, 22},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runStep1(tt.input, tt.size, tt.corrupted)
			if ans != tt.expected {
				t.Errorf(`TestStep1SampleInput got %d, expected %d`, ans, tt.expected)
			}
		})
	}
}

func TestStep2SampleInput(t *testing.T) {
	var tests = []struct {
		input     string
		size      int
		corrupted int
		expected  string
	}{
		{`5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`, 7, 12, "6,1"},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runStep2(tt.input, tt.size, tt.corrupted)
			if ans != tt.expected {
				t.Errorf(`TestStep2SampleInput got %s, expected %s`, ans, tt.expected)
			}
		})
	}
}
