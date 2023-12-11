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

const exampleInputC = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

func TestPart2C(t *testing.T) {
	answer := part2(exampleInputC)
	solution := 4
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

const exampleInputD = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

func TestPart2D(t *testing.T) {
	answer := part2(exampleInputD)
	solution := 8
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

const exampleInputE = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestPart2E(t *testing.T) {
	answer := part2(exampleInputE)
	solution := 10
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}
