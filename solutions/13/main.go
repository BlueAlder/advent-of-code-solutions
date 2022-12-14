package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func filter[T any](slice []T, testFn func(T) bool) (res []T) {
	for _, el := range slice {
		if testFn(el) {
			res = append(res, el)
		}
	}
	return res
}

func mapSlice[T any](slice []T, mapFn func(T) any) (res []any) {
	for _, el := range slice {
		res = append(res, mapFn(el))
	}
	return res
}

func parseInput(fileName string) []any {
	buf, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Unable to open file %s\n", fileName)
		fmt.Println(err)
		os.Exit(1)
	}
	input := string(buf)
	inputSplit := strings.Split(input, "\n")
	filterFn := func(e string) bool {
		return e != ""
	}
	mapFn := func(e string) any {
		var res any
		json.Unmarshal([]byte(e), &res)
		return res
	}
	filteredInput := filter(inputSplit, filterFn)
	mappedInput := mapSlice(filteredInput, mapFn)
	return mappedInput
}

func comparePackets(left []any, right []any) bool {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	iter := max(len(left), len(right))

	for i := 0; i < iter; i++ {
		if i > len(left) || i > len(right) {
			if i > len(left) {
				return true
			}
			return false
		}

	}
}

func solve(inputFileName string, part int) {
	input := parseInput(inputFileName)

	a1 := [][]any{{2}}
	a2 := [][]any{{6}}
	input = append(input, a1)
	input = append(input, a2)

}

func main() {
	solve("input.txt", 2)
}
