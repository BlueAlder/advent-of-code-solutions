package day17

import "testing"

const exampleInput string = ""

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 102
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := -1
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
