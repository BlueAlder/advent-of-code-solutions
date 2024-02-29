package day09

import "testing"

const exampleInput string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 114
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := 2
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
