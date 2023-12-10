// Solution for day10 of the Advent of Code Challenge 2023
package day10

import (
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	return part1(input)
	// if part == 1 {
	// 	return part1(input)
	// } else if part == 2 {
	// 	return part2(input)
	// } else {
	// 	util.LogFatal("invalid part number")
	// 	return -1
	// }
}

type Direction struct {
	dx int
	dy int
}

type Point struct {
	x int
	y int
}

var down Direction = Direction{0, 1}
var up Direction = Direction{0, -1}
var left Direction = Direction{-1, 0}
var right Direction = Direction{1, 0}

var directions = []Direction{up, right, down, left}

type Pipe struct {
	entries []Direction
	exits   []Direction
}

var pipes = map[string]Pipe{
	"|": {entries: []Direction{up, down}, exits: []Direction{up, down}},
	"-": {entries: []Direction{left, right}, exits: []Direction{left, right}},
	"L": {entries: []Direction{down, left}, exits: []Direction{right, up}},
	"J": {entries: []Direction{down, right}, exits: []Direction{left, up}},
	"7": {entries: []Direction{right, up}, exits: []Direction{down, left}},
	"F": {entries: []Direction{left, up}, exits: []Direction{down, right}},
}

func (p Pipe) getExitFromEntry(d Direction) (Direction, error) {
	for idx, entryDir := range p.entries {
		if entryDir == d {
			return p.exits[idx], nil
		}
	}
	return Direction{0, 0}, errors.New("Invalid Entry")
}

func part1(inputData string) int {

	lines := strings.Split(inputData, "\n")

	// Find S
	var s Point
	for y, line := range lines {
		if x := strings.Index(line, "S"); x != -1 {
			s = Point{x, y}
			break
		}

	}
	fmt.Println(s)
	curr := s
	var direction Direction
	// Find start
	for _, dir := range directions {
		sym := string(lines[curr.y+dir.dy][curr.x+dir.dx])
		pipe, exists := pipes[sym]
		if !exists {
			continue
		}
		if _, err := pipe.getExitFromEntry(dir); err == nil {
			direction = dir
			break
		}
	}

	steps := 0
	for {
		steps++
		curr.x += direction.dx
		curr.y += direction.dy
		if curr == s {
			break
		}

		sym := string(lines[curr.y][curr.x])

		pipe, exists := pipes[sym]
		if !exists {
			panic("invalid pipe direction")
		}
		var err error
		direction, err = pipe.getExitFromEntry(direction)
		if err != nil {
			panic("invalid direction")
		}
	}
	return steps / 2
}

func part2(inputData string) int {
	return 0
}
