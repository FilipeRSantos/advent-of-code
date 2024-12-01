package main

import "testing"

func TestStep1SampleInput(t *testing.T) {
	input :=
		`3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 11
	count := runStep1(input)

	if count != expected {
		t.Fatalf(`TestStep1SampleInput got %d, expected %d`, count, expected)
	}
}

func TestStep2SampleInput(t *testing.T) {
	input :=
		`3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 31
	count := runStep2(input)

	if count != expected {
		t.Fatalf(`TestStep2SampleInput got %d, expected %d`, count, expected)
	}
}
