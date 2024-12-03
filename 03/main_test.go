package main

import "testing"

func TestStep1SampleInput(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	expected := 161
	count := runStep1(input)

	if count != expected {
		t.Fatalf(`TestStep1SampleInput got %d, expected %d`, count, expected)
	}
}

func TestStep2SampleInput(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := 48
	count := runStep2(input)

	if count != expected {
		t.Fatalf(`TestStep2SampleInput got %d, expected %d`, count, expected)
	}
}
