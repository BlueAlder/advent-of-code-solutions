// Solution for day15 of the Advent of Code Challenge 2024
package day15

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

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

type cell int

const (
	empty cell = iota
	wall
	robot
	box
)

type grid struct {
	cells         [][]cell
	robotLocation image.Point
	directions    []direction
}

func parseInput(input string) *grid {
	parts := strings.Split(input, "\n\n")

	grid := &grid{
		cells:      make([][]cell, 0),
		directions: make([]direction, 0),
	}
	gridRows := strings.Split(parts[0], "\n")
	for y, row := range gridRows {
		grid.cells = append(grid.cells, make([]cell, len(row)))
		for x, c := range row {
			switch c {
			case '#':
				grid.cells[y][x] = wall
			case '.':
				grid.cells[y][x] = empty
			case '@':
				grid.cells[y][x] = empty
				grid.robotLocation = image.Pt(x, y)
			case 'O':
				grid.cells[y][x] = box
			}
		}
	}

	dirChars := strings.Join(strings.Split(parts[1], "\n"), "")
	for _, c := range dirChars {
		switch c {
		case '^':
			grid.directions = append(grid.directions, direction{0, -1})
		case '>':
			grid.directions = append(grid.directions, direction{1, 0})
		case '<':
			grid.directions = append(grid.directions, direction{-1, 0})
		case 'v':
			grid.directions = append(grid.directions, direction{0, 1})
		}
	}
	return grid
}

func part1(inputData string) int {
	grid := parseInput(inputData)

	for _, dir := range grid.directions {

		// grid.printGrid()
		// fmt.Println(dir)
		next := grid.robotLocation.Add(image.Pt(dir.dx, dir.dy))
		switch grid.cells[next.Y][next.X] {
		case wall:
			continue
		case empty:
			grid.robotLocation = next
		case box:
			boxesToMove := grid.moveBoxesIfPossible(next, dir)
			if boxesToMove == 0 {
				continue
			}
			grid.cells[next.Y][next.X] = empty
			grid.cells[next.Y+(dir.dy*boxesToMove)][next.X+(dir.dx*boxesToMove)] = box
			grid.robotLocation = next
		}
	}

	return grid.boxGPSCoordinates()
}

func part2(inputData string) int {
	return 0
}

type direction struct {
	dx int
	dy int
}

func (g grid) moveBoxesIfPossible(boxToPush image.Point, direction direction) int {
	boxesToMove := 1
	next := boxToPush.Add(image.Pt(direction.dx, direction.dy))
	for {
		if g.cells[next.Y][next.X] == wall {
			return 0
		}
		if g.cells[next.Y][next.X] == box {
			boxesToMove++
		}
		if g.cells[next.Y][next.X] == empty {
			return boxesToMove
		}
		next = next.Add(image.Pt(direction.dx, direction.dy))
	}
}

func (g *grid) boxGPSCoordinates() int {
	total := 0
	for y, row := range g.cells {
		for x, c := range row {
			if c == box {
				total += (100 * y) + x
			}
		}
	}
	return total
}

func (g *grid) printGrid() {
	for y, row := range g.cells {
		for x, c := range row {
			if g.robotLocation.X == x && g.robotLocation.Y == y {
				fmt.Print("@")
				continue
			}
			switch c {
			case wall:
				fmt.Print("#")
			case empty:
				fmt.Print(".")
			case box:
				fmt.Print("O")
			}
		}
		fmt.Print("\n")
	}
}
