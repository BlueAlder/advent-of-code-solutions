// Solution for day10 of the Advent of Code Challenge 2023
package day10

import (
	_ "embed"
	"errors"
	"reflect"
	"slices"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return part2(input)
	} else {
		util.LogFatal("invalid part number")
		return -1
	}
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
	"-": {entries: []Direction{right, left}, exits: []Direction{right, left}},
	"L": {entries: []Direction{left, down}, exits: []Direction{up, right}},
	"J": {entries: []Direction{right, down}, exits: []Direction{up, left}},
	"7": {entries: []Direction{right, up}, exits: []Direction{down, left}},
	"F": {entries: []Direction{up, left}, exits: []Direction{right, down}},
}

func (p Pipe) getExitFromEntry(d Direction) (Direction, error) {
	for idx, entryDir := range p.entries {
		if entryDir == d {
			return p.exits[idx], nil
		}
	}
	return Direction{0, 0}, errors.New("invalid Entry")
}

func part1(inputData string) int {

	lines := strings.Split(inputData, "\n")
	path := walkThePath(lines)
	return len(path) / 2
}

func getStartingPipe(s Point, lines []string) string {
	var outputs []Direction
	for _, dir := range directions {
		if !isInBounds(lines, Point{s.x + dir.dx, s.y + dir.dy}) {
			continue
		}
		sym := string(lines[s.y+dir.dy][s.x+dir.dx])
		pipe, exists := pipes[sym]
		if !exists {
			continue
		}
		if _, err := pipe.getExitFromEntry(dir); err == nil {
			outputs = append(outputs, dir)
		}
	}

	for k, v := range pipes {
		if reflect.DeepEqual(v.exits, outputs) {
			return k
		}
	}
	panic("no pipe found")
}

func findStartPoint(lines []string) Point {
	var s Point
	for y, line := range lines {
		if x := strings.Index(line, "S"); x != -1 {
			s = Point{x, y}
			break
		}
	}
	return s
}

func walkThePath(lines []string) []Point {
	var path []Point
	start := findStartPoint(lines)
	sPipe := getStartingPipe(start, lines)
	lines[start.y] = strings.Replace(lines[start.y], "S", sPipe, 1)
	dir := pipes[sPipe].exits[0]
	curr := start
	for {
		curr.x += dir.dx
		curr.y += dir.dy
		path = append(path, curr)
		if curr == start {
			break
		}

		sym := string(lines[curr.y][curr.x])

		pipe, exists := pipes[sym]
		if !exists {
			panic("invalid pipe direction")
		}
		var err error
		dir, err = pipe.getExitFromEntry(dir)
		if err != nil {
			panic("invalid direction")
		}
	}
	return path
}

func part2(inputData string) int {
	lines := strings.Split(inputData, "\n")
	// Find S

	path := walkThePath(lines)
	area := 0
	for y, line := range lines {
		crosses := 0
		for x := range line {
			p := Point{x: x, y: y}
			// on the path
			if !slices.Contains(path, p) {

				if crosses%2 != 0 {
					area++
				}
			} else {
				if slices.Contains([]string{"|", "J", "7"}, string(lines[p.y][p.x])) {
					crosses++
				}
			}
		}
	}

	return 0

}

func isInBounds(lines []string, p Point) bool {
	xInBound := p.x >= 0 && p.x < len(lines[0])
	yInBound := p.y >= 0 && p.y < len(lines)
	return xInBound && yInBound
}
