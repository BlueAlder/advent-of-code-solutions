// Solution for day17 of the Advent of Code Challenge 2023
package day17

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/BlueAlder/advent-of-code-solutions/pkg/defaultdict"
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

type Node struct {
	x           int
	y           int
	direction   int
	consecutive int
}

type Graph []string

type direction struct {
	dx int
	dy int
}

var dirs []direction = []direction{
	{dx: 1, dy: 0},
	{dx: 0, dy: 1},
	{dx: -1, dy: 0},
	{dx: 0, dy: -1},
}

var distances = defaultdict.NewDefaultDict[Node, int](math.MaxInt)

func part1(inputData string) int {
	var g Graph = strings.Split(inputData, "\n")
	return g.dijkstra(g.getNextNodes, 0)
}

func part2(inputData string) int {
	var g Graph = strings.Split(inputData, "\n")
	return g.dijkstra(g.getNextNodesMinMax, 4)
}

func (g Graph) dijkstra(neighbourFn func(Node) []Node, minDist int) int {
	startNodeUp := Node{
		x:           0,
		y:           0,
		direction:   0,
		consecutive: 0,
	}
	startNodeDown := Node{
		x:           0,
		y:           0,
		direction:   1,
		consecutive: 0,
	}

	distances.Add(startNodeUp, 0)
	distances.Add(startNodeDown, 0)

	visited := make(sets.Set[Node])
	via := make(map[Node]Node)

	q := make(PriorityQueue, 0)
	heap.Push(&q, startNodeUp)
	heap.Push(&q, startNodeDown)

	for len(q) > 0 {
		n := heap.Pop(&q).(Node)
		visited.Add(n)
		neighbours := neighbourFn(n)

		for _, neigh := range neighbours {
			if !visited.Has(neigh) && distances.Get(neigh) > distances.Get(n)+g.getHeat(neigh) {
				distances.Add(neigh, distances.Get(n)+g.getHeat(neigh))
				via[neigh] = n
				heap.Push(&q, neigh)
			}
		}
	}

	min := math.MaxInt
	var n Node
	for k, v := range distances.Values() {
		if k.y == len(g)-1 && k.x == len(g[0])-1 && k.consecutive >= minDist && v < min {
			min = v
			n = k
		}
	}

	var path []Node
	path = append(path, n)
	for {
		var ok bool
		n, ok = via[n]
		if !ok {
			break
		}
		path = append(path, n)
	}
	g.printPath(path)
	return min
}

func (g Graph) printPath(path []Node) {
	for y, row := range g {
		line := ""
		for x := range row {
			if slices.ContainsFunc(path, func(n Node) bool { return n.x == x && n.y == y }) {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func (g Graph) getNextNodes(n Node) []Node {
	ns := make([]Node, 0)
	for i, dir := range dirs {
		// Can't go backwards
		if (i+2)%len(dirs) == n.direction {
			continue
		}
		nn := Node{x: n.x + dir.dx, y: n.y + dir.dy, direction: i, consecutive: 1}
		if !g.isInBounds(nn) {
			continue
		}
		if n.direction == i {
			nn.consecutive = n.consecutive + 1
			if nn.consecutive > 3 {
				continue
			}
		}
		ns = append(ns, nn)
	}
	return ns
}

func (g Graph) getNextNodesMinMax(n Node) []Node {
	ns := make([]Node, 0)
	if n.consecutive < 4 {
		dir := dirs[n.direction]
		nn := Node{x: n.x + dir.dx, y: n.y + dir.dy, direction: n.direction, consecutive: n.consecutive + 1}
		if g.isInBounds(nn) {
			return append(ns, nn)
		}
		return ns
	}
	for i, dir := range dirs {
		// Can't go backwards
		if (i+2)%len(dirs) == n.direction {
			continue
		}
		nn := Node{x: n.x + dir.dx, y: n.y + dir.dy, direction: i, consecutive: 1}
		if !g.isInBounds(nn) {
			continue
		}
		if n.direction == i {
			nn.consecutive = n.consecutive + 1
			if nn.consecutive > 10 {
				continue
			}
		}
		ns = append(ns, nn)
	}
	return ns
}

func (g Graph) getHeat(n Node) int {
	num, err := strconv.Atoi(string(g[n.y][n.x]))
	if err != nil {
		panic("unable to convert number")
	}
	return num
}

func (g Graph) isInBounds(n Node) bool {
	return n.x >= 0 && n.x < len(g[0]) && n.y >= 0 && n.y < len(g)
}
