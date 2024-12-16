package day12

import (
	_ "embed"
	"testing"
)

//go:embed example_input.txt
var exampleInput string

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 1930
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
func TestPart1a(t *testing.T) {
	region := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	answer := part1(region)
	solution := 772
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
func TestPart1b(t *testing.T) {
	region := `AAAA
BBCD
BBCC
EEEC`
	answer := part1(region)
	solution := 140
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := 1206
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
