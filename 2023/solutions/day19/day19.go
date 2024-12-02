// Solution for day19 of the Advent of Code Challenge 2023
package day19

import (
	_ "embed"
	"regexp"
	"strconv"
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

var digitsRE = regexp.MustCompile(`\d+`)

func part1(inputData string) int {
	parts := strings.Split(inputData, "\n\n")
	wfs := parseRules(parts[0])

	prs := strings.Split(parts[1], "\n")
	total := 0
	for _, pr := range prs {
		digits, _ := util.MapSliceWithError(digitsRE.FindAllString(pr, -1), strconv.Atoi)
		p := map[string]int{
			"x": digits[0],
			"m": digits[1],
			"a": digits[2],
			"s": digits[3],
		}

		wfL := "in"
	WorkFlow:
		for {
			if wfL == "A" {
				for _, v := range p {
					total += v
				}
				break
			} else if wfL == "R" {
				break
			}

			wf := wfs[wfL]
			for _, rule := range wf.rules {
				if rule.operator(p[rule.part], rule.val) {
					wfL = rule.next
					continue WorkFlow
				}
			}
			wfL = wf.final
		}

	}
	return total
}

type Interval struct {
	min int
	max int
}

func (i Interval) getLength() int {
	return i.max - i.min + 1
}

func part2(inputData string) int {
	parts := strings.Split(inputData, "\n\n")
	wfs := parseRules(parts[0])
	defaultInterval := Interval{
		min: 1,
		max: 4000,
	}
	intervals := []Interval{defaultInterval, defaultInterval, defaultInterval, defaultInterval}
	return checkPossibilties("in", wfs, intervals)

}

var indexToPart = map[string]int{
	"x": 0,
	"m": 1,
	"a": 2,
	"s": 3,
}

func checkPossibilties(wfL string, wfs map[string]Workflow, intervals []Interval) (count int) {
	if wfL == "A" {
		x := intervals[0].getLength()
		m := intervals[1].getLength()
		a := intervals[2].getLength()
		s := intervals[3].getLength()
		return x * m * a * s
	} else if wfL == "R" {
		return 0
	}

	wf := wfs[wfL]
	for _, rule := range wf.rules {
		i := indexToPart[rule.part]
		nI := make([]Interval, 4)
		copy(nI, intervals)
		if rule.operatorSym == ">" && intervals[i].max > rule.val {
			if intervals[i].min < rule.val {
				nI[i].min = rule.val + 1
				intervals[i].max = rule.val
			}
			count += checkPossibilties(rule.next, wfs, nI)

		} else if rule.operatorSym == "<" && intervals[i].min < rule.val {
			if intervals[i].max > rule.val {
				nI[i].max = rule.val - 1
				intervals[i].min = rule.val
			}
			count += checkPossibilties(rule.next, wfs, nI)
		}

	}

	count += checkPossibilties(wf.final, wfs, intervals)
	return
}

type Workflow struct {
	rules []Rule
	final string
}

type Rule struct {
	part        string
	operator    func(int, int) bool
	operatorSym string
	val         int
	next        string
}

var wfRE = regexp.MustCompile(`[^{}]+`)

func parseRules(input string) map[string]Workflow {
	wfs := make(map[string]Workflow)
	workflows := strings.Split(input, "\n")
	for _, workflow := range workflows {
		p := wfRE.FindAllString(workflow, -1)
		label := p[0]

		rulesString := strings.Split(p[1], ",")
		final := rulesString[len(rulesString)-1]
		rulesString = rulesString[:len(rulesString)-1]

		var rules []Rule
		for _, rule := range rulesString {
			part := string(rule[0])
			operatorSym := string(rule[1])
			operater := lessThan
			if operatorSym == ">" {
				operater = greaterThan
			}
			val, _ := strconv.Atoi(digitsRE.FindString(rule))
			next := strings.Split(rule, ":")[1]

			rules = append(rules, Rule{
				part:        part,
				operator:    operater,
				operatorSym: operatorSym,
				val:         val,
				next:        next,
			})
		}
		wf := Workflow{
			rules: rules,
			final: final,
		}
		wfs[label] = wf
	}
	return wfs
}

func lessThan(a, b int) bool {
	return a < b
}

func greaterThan(a, b int) bool {
	return a > b
}
