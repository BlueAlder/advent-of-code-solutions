// Solution for day08 of the Advent of Code Challenge 2023
package day08

import (
	_ "embed"
	"regexp"
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

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value string
}

func (t *TreeNode) getLastChar() rune {
	return rune(t.value[len(t.value)-1])
}

func (t *TreeNode) getDistanceToZ(instructions string) (steps int) {
	currNode := t
	for {
		for _, char := range instructions {
			steps += 1
			if char == 'R' {
				currNode = currNode.right
			} else if char == 'L' {
				currNode = currNode.left
			}
			if currNode.getLastChar() == 'Z' {
				return
			}
		}
	}
}

type NodeLookup map[string]*TreeNode

func (nl NodeLookup) getOrCreateNode(val string) *TreeNode {
	n, exists := nl[val]
	if !exists {
		n = &TreeNode{value: val}
		nl[val] = n
	}
	return n
}

func part1(inputData string) int {

	sections := strings.Split(inputData, "\n\n")
	instructions := sections[0]

	nodeLookup, _ := buildNodeTree(sections[1])

	currNode := nodeLookup["AAA"]
	steps := 0
MainLoop:
	for {
		for _, char := range instructions {
			if char == 'R' {
				currNode = currNode.right
			} else if char == 'L' {
				currNode = currNode.left
			}
			steps++
			if currNode.value == "ZZZ" {
				break MainLoop
			}
		}
	}

	return steps
}

func part2(inputData string) int {
	sections := strings.Split(inputData, "\n\n")
	instructions := sections[0]
	_, startingNodes := buildNodeTree(sections[1])

	dists := make([]int, len(startingNodes))
	for i, start := range startingNodes {
		dists[i] = start.getDistanceToZ(instructions)
	}
	return util.LCM(dists[0], dists[1], dists[2:]...)
}

func buildNodeTree(input string) (NodeLookup, []*TreeNode) {
	lines := strings.Split(input, "\n")

	nodeLookup := make(NodeLookup)
	var startingNodes []*TreeNode
	reWords := regexp.MustCompile(`[A-Z\d]+`)

	for _, line := range lines {
		vals := reWords.FindAllString(line, -1)
		root := nodeLookup.getOrCreateNode(vals[0])
		left := nodeLookup.getOrCreateNode(vals[1])
		right := nodeLookup.getOrCreateNode(vals[2])

		root.left = left
		root.right = right
		if root.getLastChar() == 'A' {
			startingNodes = append(startingNodes, root)
		}
	}

	return nodeLookup, startingNodes
}
