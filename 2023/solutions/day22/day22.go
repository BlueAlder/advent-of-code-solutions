// Solution for day22 of the Advent of Code Challenge 2023
package day22

import (
	_ "embed"
	"slices"
	"strconv"
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
	x, y, z int
}

type Grid struct {
	bricks       []Brick
	brickMapping map[Point]Brick
}

type Brick2Brick map[Brick]sets.Set[Brick]

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
	g := parseInput(lines)
	g.LetFall()
	supporting, supportedBy := g.findSupports()

	total := 0
	for _, supported := range supporting {
		if len(supported) == 0 ||
			util.SliceEvery(supported.ToSlice(), func(b Brick) bool { return len(supportedBy[b]) > 1 }) {

			total++
			continue
		}

	}
	return total
}

func part2(inputData string) int {
	lines := strings.Split(inputData, "\n")
	g := parseInput(lines)
	g.LetFall()

	supporting, supportedBy := g.findSupports()

	count := 0
	for _, brick := range g.bricks {
		falling := make(sets.Set[Brick])
		falling.Add(brick)
		queue := []Brick{brick}

		for len(queue) > 0 {
			b := queue[0]
			queue = queue[1:]

			supported := supporting[b]
			if len(supported) > 0 {
				for _, s := range supported.ToSlice() {
					if util.SliceEvery(supportedBy[s].ToSlice(), func(sb Brick) bool { return falling.Has(sb) }) {
						if !falling.Has(s) {
							falling.Add(s)
							queue = append(queue, s)
							count++

						}

					}
				}
			}
		}

	}

	return count
}

func parseInput(lines []string) Grid {
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

	return Grid{bricks, brickMapping}
}

func (g Grid) LetFall() {
	for i, b := range g.bricks {
		if b.minZ() <= 1 {
			continue
		}
		for _, v := range b.getOccupiedPoints() {
			delete(g.brickMapping, v)
		}

		newBrick := b

	FallingLoop:
		for newBrick.minZ() > 1 {
			below := newBrick.getPointsBelow()

			for _, p := range below {
				// Hit another block
				if _, ok := g.brickMapping[p]; ok {
					break FallingLoop
				}
			}

			newBrick.end.z--
			newBrick.start.z--
		}
		for _, v := range newBrick.getOccupiedPoints() {
			g.brickMapping[v] = newBrick
		}
		g.bricks[i] = newBrick
	}

	slices.SortStableFunc(g.bricks, func(a, b Brick) int {
		return a.minZ() - b.minZ()
	})
}

func (g Grid) findSupports() (supporting Brick2Brick, supportedBy Brick2Brick) {
	supporting = make(Brick2Brick)
	supportedBy = make(Brick2Brick)

	for _, block := range g.bricks {
		supporting[block] = make(sets.Set[Brick], 0)
		supportedBy[block] = make(sets.Set[Brick], 0)
	}

	for _, block := range g.bricks {
		ps := block.getOccupiedPoints()
		for _, p := range ps {
			p.z++
			if b, ok := g.brickMapping[p]; ok && b != block {
				supporting[block].Add(b)
				supportedBy[b].Add(block)
			}
		}
	}
	return
}
