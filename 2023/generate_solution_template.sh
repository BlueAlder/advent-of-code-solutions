#!/bin/bash

cwd=$(basename "$PWD")

if [ $cwd != "2023" ]
then
  echo "Please run this from the root dir for 2023"
  exit 1
fi

if [ $# != 1 ]
then
  echo "Please provide 1 argument which is the day number of the challenge"
  exit 1
fi

daynum=$(printf "day%02d\n" $1)

cd solutions

if [ -d $daynum ]
then
  echo "Challenge Directory already exists"
  exit 1
fi

mkdir $daynum

cd $daynum

cat <<EOF > $daynum.go
// Solution for $daynum of the Advent of Code Challenge 2023
package $daynum

import (
	_ "embed"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	return part1(input)
	// if part == 1 {
	// 	return part1(input)
	// } else if part == 2 {
	// 	return part2(input)
	// } else {
	// 	util.LogFatal("invalid part number")
	// 	return -1
	// }
}

func part1(inputData string) int {
	return -1
}

func part2(inputData string) int {
	return -1
}

EOF

cat <<EOF > ${daynum}_test.go
package $daynum

import "testing"

const exampleInput string = ""

func TestPart1(t *testing.T) {
	answer := part1(exampleInput)
	solution := -1
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

func TestPart2(t *testing.T) {
	answer := part2(exampleInput)
	solution := -1
	if answer != solution {
		t.Fatalf("Example input failed. Got: %d, Want: %d", answer, solution)
	}
}

EOF


touch input.txt

echo "Generated files complete, make sure to update the function mapping in solutions.go"