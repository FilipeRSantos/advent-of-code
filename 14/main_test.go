package main

import "testing"

func TestStep1SampleInput(t *testing.T) {
	sampleInput := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	expected := 12
	count := runStep1(sampleInput, 11, 7)

	if count != expected {
		t.Fatalf(`TestStep1SampleInput got %d, expected %d`, count, expected)
	}
}
