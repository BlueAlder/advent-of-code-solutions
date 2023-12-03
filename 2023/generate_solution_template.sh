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

func Solve() {
	panic("Day $1 not implemented")
}
EOF

touch input.txt

echo "Generated files complete, make sure to update the function mapping in solutions.go"