// Solution for day14 of the Advent of Code Challenge 2023
package day14

import (
	_ "embed"
	"fmt"
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

type Grid [][]int

const (
	empty      = 0
	fixedRock  = iota
	movingRock = iota
)

func part1(inputData string) int {
	grid := parseInput(inputData)
	grid.tiltNorth()
	grid.print()

	return grid.calculateLoad()
}

func part2(inputData string) int {
	grid := parseInput(inputData)

	seen := make(map[string]int)
	seen[(grid.ToString())] = 0
	offset, length := grid.getCycleOffsetAndLength()
	numCycles := 1_000_000_000
	remaining := (numCycles - offset) % length
	grid.cycle(remaining)

	return grid.calculateLoad()
}

func parseInput(inputData string) Grid {
	rows := strings.Split(inputData, "\n")
	grid := make(Grid, len(rows))

	for y, row := range rows {
		grid[y] = make([]int, len(row))
		for x, chr := range row {
			if chr == 'O' {
				grid[y][x] = movingRock
			} else if chr == '#' {
				grid[y][x] = fixedRock
			} else {
				grid[y][x] = empty
			}
		}
	}
	return grid
}

func (g Grid) cycle(n int) {
	for i := 0; i < n; i++ {
		g.tiltNorth()
		g.tiltWest()
		g.tiltSouth()
		g.tiltEast()
	}
}

func (g Grid) getCycleOffsetAndLength() (offset int, length int) {
	seen := make(map[string]int)
	seen[(g.ToString())] = 0
	cycles := 0
	for {
		cycles++
		g.cycle(1)

		gString := g.ToString()
		if v, ok := seen[gString]; ok {
			return v, (cycles - v)
		} else {
			seen[g.ToString()] = cycles
		}
	}
}

func (g Grid) tiltNorth() {
	for y, row := range g {
		for x, val := range row {
			if val == movingRock {
				lowestY := y
				for newY := y - 1; newY >= 0; newY-- {
					if g[newY][x] == empty {
						lowestY = newY
						continue
					}
					break
				}
				if lowestY != y {
					g[y][x] = empty
					g[lowestY][x] = movingRock
				}
			}
		}
	}
}

func (g Grid) tiltWest() {
	for y, row := range g {
		for x, val := range row {
			if val == movingRock {
				lowestX := x
				for newX := x - 1; newX >= 0; newX-- {
					if g[y][newX] == empty {
						lowestX = newX
						continue
					}
					break
				}
				if lowestX != x {
					g[y][x] = empty
					g[y][lowestX] = movingRock
				}
			}
		}
	}
}

func (g Grid) tiltSouth() {
	for y := len(g) - 1; y >= 0; y-- {
		for x, val := range g[y] {
			if val == movingRock {
				highestY := y
				for newY := y + 1; newY < len(g); newY++ {
					if g[newY][x] == empty {
						highestY = newY
						continue
					}
					break
				}
				if highestY != y {
					g[y][x] = empty
					g[highestY][x] = movingRock
				}
			}
		}
	}
}

func (g Grid) tiltEast() {
	for y, row := range g {
		for x := len(row) - 1; x >= 0; x-- {
			if g[y][x] == movingRock {

				highestX := x
				for newX := x + 1; newX < len(row); newX++ {
					if g[y][newX] == empty {
						highestX = newX
						continue
					}
					break
				}
				if highestX != x {
					g[y][x] = empty
					g[y][highestX] = movingRock
				}
			}
		}
	}
}

func (g Grid) calculateLoad() (total int) {
	for y, row := range g {
		for _, val := range row {
			if val == movingRock {
				total += len(g) - y
			}
		}
	}
	return
}

func (g Grid) ToString() string {
	line := ""
	for _, row := range g {
		for _, val := range row {
			if val == movingRock {
				line += "O"
			} else if val == fixedRock {
				line += "#"
			} else {
				line += "."
			}
		}
		line += "\n"
	}
	return line
}

func (g Grid) print() {
	fmt.Println(g.ToString())
}
