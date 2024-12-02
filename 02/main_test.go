package main

import "testing"

var sampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestStep1SampleInput(t *testing.T) {

	expected := 2
	count := runStep1(sampleInput)

	if count != expected {
		t.Fatalf(`TestStep1SampleInput got %d, expected %d`, count, expected)
	}
}

func TestStep2SampleInput(t *testing.T) {
	expected := 4
	count := runStep2(sampleInput)

	if count != expected {
		t.Fatalf(`TestStep2SampleInput got %d, expected %d`, count, expected)
	}
}
