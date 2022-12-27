#!/usr/bin/env python3
# Advent of Code Challenge 23

import os
import functools
import operator
from collections import defaultdict


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
  
  def moveToProposal(self) -> None:
    self.x = self.proposal[0]
    self.y = self.proposal[1]
    self.proposal = None
  
  def __str__(self) -> str:
    return f"({self.x}, {self.y}), p:{self.proposal}"

def parseInput(filename):
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for y, line in enumerate(f):
      for x, chr in enumerate(line):
        if chr == "#":
          elves.append(Elf(x, y))

def getAdjacentPositions(p):
  # not finsihed
  adjP = []
  for cardinal in position_check:
    adjP.append(list(map(lambda x: (p[0] + x[0], p[1] + x[1]), cardinal)))
  return adjP

def getMinMaxCoords():
  xCoords = list(map(lambda e: e.x, elves))
  yCoords = list(map(lambda e: e.y, elves))
  minx = min(xCoords)
  maxx = max(xCoords)
  miny = min(yCoords)
  maxy = max(yCoords)
  return minx, maxx, miny, maxy

def getEmptyNumberTiles():
  minx, maxx, miny, maxy = getMinMaxCoords()
  return (maxx - minx + 1) * (maxy - miny + 1) - len(elves)

def playRound(round_number):
  elfPositions = set()
  proposals = defaultdict(lambda: 0)

  for elf in elves:
    elfPositions.add((elf.x, elf.y))
  
  for elf in elves:
    ap = getAdjacentPositions((elf.x, elf.y))
    uniqueAp = set(functools.reduce(operator.iconcat, ap, []))
    adjElves = elfPositions.intersection(uniqueAp)
    if len(adjElves) > 0:
      for i in range(4):
        if adjElves.isdisjoint(set(ap[(i + round_number) % 4])):
          elf.proposal = ap[(i + round_number) % 4][0]
          proposals[ap[(i + round_number) % 4][0]] += 1
          break
  
  for i, elf in enumerate(elves):
    if proposals[elf.proposal] == 1:
      elf.moveToProposal()
    elf.proposal = None
    # print(i, elf)


    
    


def solve(filename, rounds):
  parseInput(filename)
  # print(getEmptyNumberTiles())
  for i in range(rounds):
    playRound(i)
  return getEmptyNumberTiles()
  
  # for e in elves:
    # print(e)
  # print(len(elves))

def main():
  p1 = solve("input.txt", 10)
  print("Part 1:", p1)

if __name__ == "__main__":
  main()
