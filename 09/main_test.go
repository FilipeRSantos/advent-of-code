package main

import (
	"testing"
)

var input = "2333133121414131402"

func TestStep1SampleInput(t *testing.T) {
	expected := 1928
	count := runStep1(input)

	if count != expected {
		t.Fatalf(`TestStep2SampleInput got %d, expected %d`, count, expected)
	}
}

func TestStep2SampleInput(t *testing.T) {
	expected := 2858
	count := runStep2(input)

	if count != expected {
		t.Fatalf(`TestStep2SampleInput got %d, expected %d`, count, expected)
	}
}
