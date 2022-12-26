#!/usr/bin/env python3
# Advent of Code Challenge 23

import os

elves = []
position_check = [
  # North
  [(0, -1), (1, -1), (-1, -1)],
  # South
  [(0, 1), (1, 1), (-1, 1)],
  # West
  [(-1, 0), (-1, -1), (-1, 1)],
  # East
  [(1, 0), (1, -1), (1, 1)]

]

class Elf:
  def __init__(self, x, y) -> None:
    self.x = x
    self.y = y
    self.proposal = None
  
  def __str__(self) -> str:
    return f"({self.x}, {self.y})"

def parseInput(filename):
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for y, line in enumerate(f):
      for x, chr in enumerate(line):
        if chr == "#":
          elves.append(Elf(x, y))

def getAdjacentPositions(p):
  # not finsihed
  directions = [(0, -1), (1, -1), (-1, -1), (0, 1), (1, 1), (-1, 1), (-1, 0), (1, 0) ]    
  adjP = []
  for dx, dy in directions:
    adjP.append((p[0] + dx, p[1] + dy))
  return adjP

def playRound(round_number):
  elfPositions = set()
  for elf in elves:
    elfPositions.add((elf.x, elf.y))
  
  for elf in elves:
    ap = getAdjacentPositions((elf.x, elf.y))
    adjElves = elfPositions.intersection(ap)
    if len(adjElves) > 0:
      for i in range(4):
        if adjElves.isdisjoint(position_check[(i + round_number) % 4])


    
    


def solve(filename):
  parseInput(filename)
  playRound()
  # for e in elves:
    # print(e)
  # print(len(elves))

def main():
  solve("input.txt")

if __name__ == "__main__":
  main()
