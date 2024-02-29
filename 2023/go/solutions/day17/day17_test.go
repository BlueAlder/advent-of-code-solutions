package day17

import "testing"

const exampleInput string = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 102
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := 94
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

const exampleInput2 string = `111111111111
999999999991
999999999991
999999999991
999999999991`

func TestPart2b(t *testing.T) {
	answer := part2(exampleInput2)
	solution := 71
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

const exampleInput3 string = `1111199999999999
9999199999999999
9999199999999999
9999199999999999
9999111111111111`

func TestPart2c(t *testing.T) {
	answer := part2(exampleInput3)
	solution := 51
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
