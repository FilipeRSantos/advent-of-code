package main

import "testing"

var input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestStep1SampleInput(t *testing.T) {
	expected := 18
	count := runStep1(input)

	if count != expected {
		t.Fatalf(`TestStep1SampleInput got %d, expected %d`, count, expected)
	}
}

func TestStep2SampleInput(t *testing.T) {
	expected := 48
	count := runStep2(input)

	if count != expected {
		t.Fatalf(`TestStep2SampleInput got %d, expected %d`, count, expected)
	}
}
