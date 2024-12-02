// Solution for day23 of the Advent of Code Challenge 2023
package day23

import (
	_ "embed"
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

type Point struct {
	x, y int
}

type Node struct {
	Point
	direction   direction
	connections []connection
}

type connection struct {
	distance int
	node     *Node
}

type direction struct {
	dx, dy int
}

type Grid struct {
	g []string
}

func part1(inputData string) int {
	g := parseInput(inputData)
	start := Point{x: 1, y: 0}
	startNode := &Node{
		Point:       start,
		direction:   direction{0, 1},
		connections: make([]connection, 0),
	}

	vistedNodes := make(sets.Set[Point])
	vistedNodes.Add(start)

	nodeQueue := make(util.Queue[*Node], 0)
	nodeQueue.Enqueue(startNode)

	for len(nodeQueue) > 0 {
		n, err := nodeQueue.Dequeue()
		if err != nil {
			panic(err)
		}
		vistedPoints := make(sets.Set[Point])
		vistedPoints.Add(n.Point)

		currPoint := n.Point

		for {
			next := g.getNextAvailablePoints(currPoint)
			if len(next) == 1 {
				if g.isSlope(next[0]) {
					// Create new node
					nn := &Node{
						Point: currPoint,
						// direction:   direction{},
						connections: []connection{},
					}

					n.connections = append(n.connections, connection{
						distance: len(vistedPoints),
						node:     nn,
					})
				}
			}

		}

	}

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
	return Grid{
		g: lines,
	}
}

func (g Grid) getNextAvailablePoints(p Point) []Point {
	dirs := []direction{
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

		char := rune(g.g[np.y][np.x])

		if g.isInBounds(np) && char != '#' {
			continue
		}

		if v, ok := slopeMap[char]; ok {
			if dir == v {
				adjPoints = append(adjPoints, np)
			}
		} else {
			adjPoints = append(adjPoints, np)

		}
	}
	return adjPoints
}

func (g Grid) isInBounds(p Point) bool {
	return p.x >= 0 && p.x < len(g.g[0]) && p.y >= 0 && p.y < len(g.g)
}

func (g Grid) isSlope(p Point) bool {
	if !g.isInBounds(p) {
		return false
	}
	char := rune(g.g[p.y][p.x])
	if _, ok := slopeMap[char]; ok {
		return true
	}
	return false
}
