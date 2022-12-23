#!/usr/bin/env python3
# Advent of Code Challenge 22

import os

open_tiles = set()
wall_tiles = set()
instructions = ""

def parseInput(filename):
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for yidx, line in enumerate(f):
      line = line.strip()
      if line == "": break
      y = yidx + 1 # off by one errors baby
      for xidx, chr in enumerate(line):
        x = xidx + 1
        if chr == ".":
          open_tiles.add((x, y))
        elif chr == "#":
          wall_tiles.add((x, y))
    f.readline()



def solve(filename):
  coords = parseInput(filename)

def main():
  solve("input.txt")

if __name__ == "__main__":
  main()
