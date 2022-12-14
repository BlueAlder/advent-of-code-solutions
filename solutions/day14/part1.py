#!/usr/bin/env python3
# Advent of Code Challenge 14

import itertools

def main():
  solve("input.txt")

def parseInput(fileName):
  coords = []
  with open("input.txt") as f:
    for line in f:
      rockLine = line.strip().split(" -> ")
      rockLine = list(map(lambda x: (int(x.split(",")[0]), int(x.split(",")[1])), rockLine))
      coords.append(rockLine)
  return coords
  
def createFlatArrOfCoords(coords):
  rockPairs = set()
  for coord in coords:
    for c1, c2 in list(zip( coord, coord[1:])):
      for x in range(c1[0], c2[0] + 1):
        rockPairs.add((x, c1[1]))
      for y in range(c1[1], c2[1] + 1):
        rockPairs.add((c1[0], y))
  return rockPairs

def lowestRock(rockPairs):
  max = 0
  for x,y in rockPairs:
    if y > max: max = y
  return max

def pourSand(starting, rockPairs, lowestRock):
  pass

def solve(fileName):
  coords = parseInput(fileName)
  rockPairs = createFlatArrOfCoords(coords)
  lowestRock = lowestRock(rockPairs)


  

if __name__ == "__main__":
  main()
