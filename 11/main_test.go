package main

import (
	"fmt"
	"testing"
)

func TestStepsSampleInput(t *testing.T) {
	var tests = []struct {
		input    string
		blinks   int
		expected int
	}{
		{`125 17`, 25, 55312},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runStep(tt.input, tt.blinks)
			if ans != tt.expected {
				t.Errorf(`TestStepsSampleInput got %d, expected %d`, ans, tt.expected)
			}
		})
	}
}
