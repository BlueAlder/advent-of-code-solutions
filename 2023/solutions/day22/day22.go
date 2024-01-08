// Solution for day22 of the Advent of Code Challenge 2023
package day22

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/BlueAlder/advent-of-code-solutions/pkg/sets"
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

func (b Brick) getOccupiedPoints() []Point {
	points := make([]Point, 0)
	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			for z := b.start.z; z <= b.end.z; z++ {
				points = append(points, Point{x, y, z})
			}
		}
	}
	return points
}

func (b Brick) getPointsBelow() []Point {
	points := make([]Point, 0)
	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			points = append(points, Point{x, y, b.minZ() - 1})
		}
	}
	return points
}

func part1(inputData string) int {
	lines := strings.Split(inputData, "\n")

	var bricks []Brick
	brickMapping := make(map[Point]Brick)
	highestZ := 0

	for _, line := range lines {
		ends := strings.Split(line, "~")
		startCoords, _ := util.MapSliceWithError(strings.Split(ends[0], ","), strconv.Atoi)
		sP := Point{x: startCoords[0], y: startCoords[1], z: startCoords[2]}
		endCoords, _ := util.MapSliceWithError(strings.Split(ends[1], ","), strconv.Atoi)
		eP := Point{x: endCoords[0], y: endCoords[1], z: endCoords[2]}
		b := Brick{start: sP, end: eP}
		bricks = append(bricks, b)

		// Map points to grid
		for _, p := range b.getOccupiedPoints() {
			brickMapping[p] = b
		}
		if util.Max(sP.z, eP.z) > highestZ {
			highestZ = util.Max(sP.z, eP.z)
		}
	}

	// Order bricks by lowest z
	slices.SortStableFunc(bricks, func(a, b Brick) int {
		return a.minZ() - b.minZ()
	})

	for i, b := range bricks {
		if b.minZ() <= 1 {
			continue
		}
		for _, v := range b.getOccupiedPoints() {
			delete(brickMapping, v)
		}

		newBrick := b

	FallingLoop:
		for newBrick.minZ() > 1 {
			below := newBrick.getPointsBelow()

			for _, p := range below {
				// Hit another block
				if _, ok := brickMapping[p]; ok {
					break FallingLoop
				}
			}

			newBrick.end.z--
			newBrick.start.z--
		}
		for _, v := range newBrick.getOccupiedPoints() {
			brickMapping[v] = newBrick
		}
		bricks[i] = newBrick

	}

	slices.SortStableFunc(bricks, func(a, b Brick) int {
		return a.minZ() - b.minZ()
	})

	supporting := make(map[Brick]sets.Set[Brick])
	supportedBy := make(map[Brick]sets.Set[Brick])

	for _, block := range bricks {
		supporting[block] = make(sets.Set[Brick], 0)
		supportedBy[block] = make(sets.Set[Brick], 0)
	}

	for _, block := range bricks {
		ps := block.getOccupiedPoints()
		for _, p := range ps {
			p.z++
			if b, ok := brickMapping[p]; ok && b != block {
				supporting[block].Add(b)
				supportedBy[b].Add(block)
			}
		}
	}

	total := 0
	for _, supported := range supporting {
		if len(supported) == 0 {
			total++
			continue
		}

		if util.SliceEvery(supported.ToSlice(), func(b Brick) bool { return len(supportedBy[b]) > 1 }) {
			total++
		}

	}
	return total
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
