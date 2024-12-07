package utils

import (
	"fmt"
	"testing"
)

func TestLeftPadWithZeros(t *testing.T) {
	var tests = []struct {
		input  string
		char   string
		length int
		expect string
	}{
		{"123", "0", 5, "00123"},
		{"123", "0", 2, "123"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%d", tt.input, tt.length)
		t.Run(testname, func(t *testing.T) {
			ans := LeftPadWith(tt.input, tt.char, tt.length)
			if ans != tt.expect {
				t.Errorf(`TestLeftPadWithZeros got %s, expected %s`, ans, tt.expect)
			}
		})
	}
}
