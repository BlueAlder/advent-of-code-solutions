// Solution for day10 of the Advent of Code Challenge 2023
package day10

import (
	_ "embed"
	"errors"
	"reflect"
	"strings"

	"github.com/BlueAlder/advent-of-code-solutions/common/sets"
	util "github.com/BlueAlder/advent-of-code-solutions/common/utils"
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

type PipeMap []string

func part1(inputData string) int {

	var m PipeMap = strings.Split(inputData, "\n")
	path := m.walkThePath()
	return len(path) / 2
}

func part2(inputData string) (area int) {
	var m PipeMap = strings.Split(inputData, "\n")
	upFacing := make(sets.Set[string])
	upFacing.Add("|", "J", "L")
	path := m.walkThePath()
	for y, line := range m {
		inside := false
		for x := range line {
			p := Point{x: x, y: y}
			if !path.Has(p) {
				if inside {
					area++
				}
			} else {
				// Check for UP facing pipes
				if upFacing.Has(string(m[p.y][p.x])) {
					inside = !inside
				}
			}
		}
	}
	return
}

func (m PipeMap) findStartPoint() Point {
	var s Point
	for y, line := range m {
		if x := strings.Index(line, "S"); x != -1 {
			s = Point{x, y}
			break
		}
	}
	return s
}

func (m PipeMap) getStartingPipe(s Point) string {
	var outputs []Direction
	for _, dir := range directions {
		if !m.isInBounds(Point{s.x + dir.dx, s.y + dir.dy}) {
			continue
		}
		sym := string(m[s.y+dir.dy][s.x+dir.dx])
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

func (m PipeMap) walkThePath() sets.Set[Point] {
	path := make(sets.Set[Point])
	start := m.findStartPoint()
	sPipe := m.getStartingPipe(start)
	m[start.y] = strings.Replace(m[start.y], "S", sPipe, 1)
	dir := pipes[sPipe].exits[0]
	curr := start
	for {
		curr.x += dir.dx
		curr.y += dir.dy
		path.Add(curr)
		if curr == start {
			break
		}

		sym := string(m[curr.y][curr.x])

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

func (m PipeMap) isInBounds(p Point) bool {
	xInBound := p.x >= 0 && p.x < len(m[0])
	yInBound := p.y >= 0 && p.y < len(m)
	return xInBound && yInBound
}
