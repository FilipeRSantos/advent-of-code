package main

import (
	"fmt"
	"testing"
)

func TestStep1SampleInput(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`, "4,6,3,5,6,3,5,2,1,0"},
		{`Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`, "0,1,2"},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			ans := runStep1(tt.input)
			if ans != tt.expected {
				t.Errorf(`TestStep1SampleInput got %s, expected %s`, ans, tt.expected)
			}
		})
	}
}

func TestSample1(t *testing.T) {
	input := `Register A: 0
Register B: 0
Register C: 9

Program: 2,6`
	expected := 1
	runStep1(input)
	if b != expected {
		t.Errorf(`TestSample1 got %d, expected %d`, b, expected)
	}
}

func TestSample3(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	expectedAPointer := 0
	expectedOutput := "4,2,5,6,7,7,7,7,3,1,0"
	output := runStep1(input)
	if output != expectedOutput {
		t.Errorf(`TestSample3 got %s, expected %s`, output, expectedOutput)
	}

	if a != expectedAPointer {
		t.Errorf(`TestSample3 got %d, expected %d`, a, expectedAPointer)
	}
}

func TestSample4(t *testing.T) {
	input := `Register A: 0
Register B: 29
Register C: 0

Program: 1,7`
	expected := 26
	runStep1(input)
	if b != expected {
		t.Errorf(`TestSample4 got %d, expected %d`, b, expected)
	}
}

func TestSample5(t *testing.T) {
	input := `Register A: 0
Register B: 2024
Register C: 43690

Program: 4,0`
	expected := 44354
	runStep1(input)
	if b != expected {
		t.Errorf(`TestSample5 got %d, expected %d`, b, expected)
	}
}
