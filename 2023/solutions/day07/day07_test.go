package day07

import "testing"

const exampleInput string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := 6440
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := 5905
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
