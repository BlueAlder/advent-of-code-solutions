package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Set[T comparable] map[T]struct{}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func mapToPoints(s []string) []Point {
	var p []Point
	for _, v := range s {
		coord := strings.Split(v, ",")
		x, err := strconv.Atoi(coord[0])
		y, err := strconv.Atoi(coord[1])
		if err != nil {
			fmt.Println("Unable to convert to int")
			os.Exit(1)
		}
		p = append(p, Point{x, y})
	}
	return p
}

func maxYValue(s Set[Point]) int {
	max := 0
	for p := range s {
		if p.y > max {
			max = p.y
		}
	}
	return max
}

func createFlatArrOfCoords(rockLines [][]Point) Set[Point] {
	rockCoords := make(Set[Point])
	for _, rockLine := range rockLines {
		for i := 0; i < len(rockLine)-1; i++ {
			c1 := rockLine[i]
			c2 := rockLine[i+1]
			for x := min(c1.x, c2.x); x < max(c1.x, c2.x)+1; x++ {
				rockCoords[Point{x, c1.y}] = struct{}{}
			}
			for y := min(c1.y, c2.y); y < max(c1.y, c2.y)+1; y++ {
				rockCoords[Point{c1.x, y}] = struct{}{}
			}
		}
	}
	return rockCoords
}

func parseInput(fileName string) [][]Point {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to open file %s\n", fileName)
		fmt.Println(err)
		os.Exit(1)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rockLines [][]Point
	for scanner.Scan() {
		rockLine := strings.Split(scanner.Text(), " -> ")
		rockLines = append(rockLines, mapToPoints(rockLine))
	}
	return rockLines
}

func getAdjPositions(p Point) []Point {
	var movements = []Point{
		{0, 1},
		{-1, 1},
		{1, 1},
	}
	positions := make([]Point, 3)
	for i, m := range movements {
		positions[i] = Point{p.x + m.x, p.y + m.y}
	}
	return positions
}

func getPossibleNextPositions(sandPosition Point, rocks Set[Point], floor int) []Point {
	nextPositions := getAdjPositions(sandPosition)
	var validMovements []Point
	for _, v := range nextPositions {
		if _, ok := rocks[v]; !(ok || v.y >= floor) {
			validMovements = append(validMovements, v)
		}
	}
	return validMovements
}

func pourSand(startingPoint Point, rockCoords Set[Point], floor int) int {
	visited := make(Set[Point])
	queue := make(chan Point, 5000)
	queue <- startingPoint
	visited[startingPoint] = struct{}{}

	for sandCoord := range queue {
		for _, nextPos := range getPossibleNextPositions(sandCoord, rockCoords, floor) {
			if _, ok := visited[nextPos]; !ok {
				queue <- nextPos
				visited[nextPos] = struct{}{}
			}
		}
		if len(queue) == 0 {
			break
		}
	}
	return len(visited)
}

func solve(fileName string, startingPoint Point, floorDelta int) {
	rockLines := parseInput(fileName)
	rockCoords := createFlatArrOfCoords(rockLines)
	floor := maxYValue(rockCoords)
	sandBFS := pourSand(startingPoint, rockCoords, floor)
	fmt.Println("BFS:", sandBFS)
}

func main() {
	solve("/usr/local/google/home/samcalamos/Documents/personal/advent-of-code-22/solutions/day14/input.txt", Point{500, 0}, 2)
}
