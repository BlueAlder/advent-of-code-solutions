// Solution for day23 of the Advent of Code Challenge 2023
package day23

import (
	_ "embed"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
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

type Point struct {
	x, y int
}

type Node struct {
	Point
	nextNodes []*Node
	direction direction
}

type direction [2]int

type Grid struct {
	g     []string
	nodes []*Node
}

func part1(inputData string) int {
	g := parseInput(inputData)

	// startNode := Node{Point{x: 1, y: 0}}

	return 0
}

func part2(inputData string) int {
	return 0
}

var slopeMap = map[rune]direction{
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
	'^': {0, -1},
}

func parseInput(inputData string) Grid {
	lines := strings.Split(inputData, "\n")
	nodes := make([]*Node, 0)

	for y, line := range lines {
		for x, v := range line {
			if v, ok := slopeMap[v]; ok {
				nodes = append(nodes, &Node{
					Point:     Point{x: x, y: y},
					nextNodes: make([]*Node, 0),
					direction: v,
				})
			}
		}
	}

	return Grid{
		g:     lines,
		nodes: nodes,
	}
}

func (g Grid) getNextAvailablePoints(p Point) []Point {
	dirs := []struct{ dx, dy int }{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, 1},
	}

	adjPoints := make([]Point, 0)

	for _, dir := range dirs {
		np := p
		np.x += dir.dx
		np.y += dir.dy

		if g.isInBounds(np) && g.g[np.y][np.x] != '#' {
			adjPoints = append(adjPoints, np)
		}
	}
	return adjPoints
}

func (g Grid) isInBounds(p Point) bool {
	return p.x >= 0 && p.x < len(g.g[0]) && p.y >= 0 && p.y < len(g.g)
}
