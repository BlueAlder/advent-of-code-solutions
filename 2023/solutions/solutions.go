package solutions

import (
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day01"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day02"
	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

var slns = map[int]interface{}{
	1: day01.Solve,
	2: day02.Solve,
}

func Run(day int) {
	if v, ext := slns[day]; ext {
		v.(func())()
		return
	}

	util.LogFatal("day does not exist in function map")
}
