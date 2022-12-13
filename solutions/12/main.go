package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x     int
	y     int
	value int
}

type Graph struct {
	grid      [][]Point
	width     int
	height    int
	startNode Point
	endNode   Point
}

func NewGraph(inputFileName string) *Graph {
	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Printf("Unable to open file %s\n", inputFileName)
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	y := 0
	var grid [][]Point
	var endNode Point
	var startNode Point
	for scanner.Scan() {
		row := []Point{}
		for x, chr := range scanner.Text() {
			if chr == 'S' {
				startNode = Point{x, y, 1}
				row = append(row, startNode)
			} else if chr == 'E' {
				endNode = Point{x, y, 26}
				row = append(row, endNode)
			} else {
				row = append(row, Point{x, y, int(chr) - 96})
			}
		}
		grid = append(grid, row)
		y++
	}
	return &Graph{
		grid:      grid,
		width:     len(grid[0]),
		height:    len(grid),
		endNode:   endNode,
		startNode: startNode,
	}
}

func (g *Graph) getAdjacentPoints(currentPoint Point) []Point {
	point_val := currentPoint.value
	var points []Point
	movements := []struct {
		x int
		y int
	}{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range movements {
		testX := d.x + currentPoint.x
		testY := d.y + currentPoint.y
		if isInRange(0, g.width, testX) && isInRange(0, g.height, testY) && g.grid[testY][testX].value+1 >= point_val {
			points = append(points, g.grid[testY][testX])
		}
	}
	return points
}

func (g *Graph) bfs(startNode Point, endNode Point, part int) int {
	visited := make(map[Point]struct{})
	dq := make(chan Point, 5000)
	nodeSteps := make(map[Point]int)

	getNodeCost := func(p Point) int {
		if v, ok := nodeSteps[p]; ok {
			return v
		}
		return 1e7
	}
	nodeSteps[startNode] = 0
	dq <- startNode

	for node := range dq {
		if (part == 1 && node == endNode) || part == 2 && node.value == endNode.value {
			return getNodeCost(node)
		}
		visited[node] = struct{}{}
		for _, adjPoint := range g.getAdjacentPoints(node) {
			if _, ok := visited[adjPoint]; ok {
				continue
			}
			new_cost := getNodeCost(node) + 1
			if new_cost < getNodeCost(adjPoint) {
				nodeSteps[adjPoint] = new_cost
				dq <- adjPoint
			}
		}
	}
	return -1
}

func isInRange(min int, max int, value int) bool {
	return value >= min && value < max
}

func main() {
	grid := NewGraph("input.txt")
	fmt.Println("Part 1:", grid.bfs(grid.endNode, grid.startNode, 1))
	fmt.Println("Part 2:", grid.bfs(grid.endNode, grid.startNode, 2))
}
