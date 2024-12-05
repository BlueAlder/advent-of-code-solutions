package day03

import (
	_ "embed"
	"testing"
)

//go:embed example_input.txt
var exampleInput string

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 161
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	overrideTestInput := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	answer := part2(overrideTestInput)
	solution := 48
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
