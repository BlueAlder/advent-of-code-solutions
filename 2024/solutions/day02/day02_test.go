package day02

import (
	_ "embed"
	"testing"
)

//go:embed example_input.txt
var exampleInput string

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 2
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := 4
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
