// Solution for day22 of the Advent of Code Challenge 2023
package day22

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
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
	x, y, z int
}

type Brick struct {
	start, end Point
}

func (b Brick) minZ() int {
	return util.Min(b.start.z, b.end.z)
}

func (b Brick) maxZ() int {
	return util.Max(b.start.z, b.end.z)
}

func part1(inputData string) int {
	lines := strings.Split(inputData, "\n")

	var bricks []Brick
	highestZ := 0

	for _, line := range lines {
		ends := strings.Split(line, "~")
		startCoords, _ := util.MapSliceWithError(strings.Split(ends[0], ","), strconv.Atoi)
		sP := Point{x: startCoords[0], y: startCoords[1], z: startCoords[2]}
		endCoords, _ := util.MapSliceWithError(strings.Split(ends[1], ","), strconv.Atoi)
		eP := Point{x: endCoords[0], y: endCoords[1], z: endCoords[2]}
		bricks = append(bricks, Brick{start: sP, end: eP})
		if util.Max(sP.z, eP.z) > highestZ {
			highestZ = util.Max(sP.z, eP.z)
		}
	}

	// printAxis(bricks, true)
	// fmt.Println()
	// printAxis(bricks, false)
	// fmt.Println()

	// Order bricks by lowest z
	slices.SortStableFunc(bricks, func(a, b Brick) int {
		return a.minZ() - b.minZ()
	})

	for i, b := range bricks {
		if b.minZ() <= 1 {
			continue
		}
		newBrick := b

	FallingLoop:
		for newBrick.minZ() > 1 {
			newBrick.end.z--
			newBrick.start.z--

			for _, testBrick := range bricks {
				if testBrick.minZ() > newBrick.minZ() {
					break
				}
				if bricksIntersect(testBrick, newBrick) {
					newBrick.end.z++
					newBrick.start.z++
					break FallingLoop
				}

			}
		}

		bricks[i] = newBrick
	}

	slices.SortStableFunc(bricks, func(a, b Brick) int {
		return a.minZ() - b.minZ()
	})

	// for _, b := range bricks {

	// }

	// printAxis(bricks, true)
	// fmt.Println()
	// printAxis(bricks, false)
	return 0
}

func printAxis(bricks []Brick, xAxis bool) {
	var grid [10][3]string
	for i := range grid {
		grid[i] = [3]string{"#", "#", "#"}
	}

	for i, b := range bricks {
		if xAxis {
			for x := util.Min(b.start.x, b.end.x); x <= util.Max(b.start.x, b.end.x); x++ {
				grid[b.start.z][x] = string(rune(i + 65))
			}
		} else {
			for y := util.Min(b.start.y, b.end.y); y <= util.Max(b.start.y, b.end.y); y++ {
				grid[b.start.z][y] = string(rune(i + 65))
			}
		}

		for z := b.maxZ(); z >= b.minZ(); z-- {
			if xAxis {
				grid[z][b.start.x] = string(rune(i + 65))
			} else {
				grid[z][b.start.y] = string(rune(i + 65))

			}
		}
	}
	for z := 9; z >= 0; z-- {
		fmt.Println(grid[z][0] + grid[z][1] + grid[z][2] + " " + strconv.Itoa(z))
	}
}

func bricksIntersect(a, b Brick) bool {
	return true &&
		a.start.x <= b.end.x &&
		a.end.x >= b.start.x &&
		a.start.y <= b.end.y &&
		a.end.y >= b.start.y &&
		a.start.z <= b.end.z &&
		a.end.z >= b.start.z
}

func part2(inputData string) int {
	return 0
}
