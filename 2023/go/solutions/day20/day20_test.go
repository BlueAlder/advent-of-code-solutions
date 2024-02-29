package day20

import "testing"

const exampleInput string = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

func TestPart1a(t *testing.T) {
	answer := part1(exampleInput)
	solution := 32000000
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

const exampleInput2 string = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func TestPart1b(t *testing.T) {
	answer := part1(exampleInput2)
	solution := 11687500
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
