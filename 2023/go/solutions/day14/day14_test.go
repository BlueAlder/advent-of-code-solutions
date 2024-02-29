package day14

import "testing"

const exampleInput string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 136
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := 64
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
