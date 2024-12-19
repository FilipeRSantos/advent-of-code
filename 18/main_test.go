package main

import (
	"fmt"
	"testing"
)

func TestSolveMaze(t *testing.T) {
	var tests = []struct {
		input     string
		size      int
		corrupted int
		part1     int
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
			ans1, _ := solve(tt.input, tt.size, tt.corrupted)
			if ans1 != tt.part1 {
				t.Errorf(`TestSolveMaze differed in part 1. Got %d, expected %d`, ans1, tt.part1)
			}
		})
	}
}
