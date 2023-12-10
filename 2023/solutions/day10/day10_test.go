package day10

import "testing"

const exampleInputA string = `.....
.S-7.
.|.|.
.L-J.
.....`

func TestPart1a(t *testing.T) {
	answer := part1(exampleInputA)
	solution := 4
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

const exampleInputB string = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

func TestPart1b(t *testing.T) {
	answer := part1(exampleInputB)
	solution := 8
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInputA)
	solution := -1
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
