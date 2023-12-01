#!/bin/bash

cwd=$(basename "$PWD")

if [ $cwd != "advent-of-code-22" ]
then
  echo "Please run this from the root dir"
  exit 1
fi

if [ $# != 1 ]
then
  echo "Please provide 1 argument which is the day number of the challenge"
  exit 1
fi

dirname=$(printf "day%02d\n" $1)

cd solutions

if [ -d $dirname ]
then
  echo "Challenge Directory already exists"
  exit 1
fi

mkdir $dirname

cd $dirname

cat <<EOF > solve.py
#!/usr/bin/env python3
# Advent of Code Challenge $1

import os

def parseInput(filename):
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    pass

def solve(filename):
  pass

def main():
  pass

if __name__ == "__main__":
  main()
EOF

chmod +x solve.py 

touch input.txt
