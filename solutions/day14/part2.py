#!/usr/bin/env python3
# Advent of Code Challenge 14

import itertools

def main():
  solve("/home/sam/Documents/advent-of-code-22/solutions/day14/input.txt")

def parseInput(fileName):
  coords = []
  with open(fileName) as f:
    for line in f:
      rockLine = line.strip().split(" -> ")
      rockLine = list(map(lambda x: (int(x.split(",")[0]), int(x.split(",")[1])), rockLine))
      coords.append(rockLine)
  return coords
  
def createFlatArrOfCoords(coords):
  rockCoords = set()
  for coord in coords:
    for c1, c2 in list(zip( coord, coord[1:])):
      for x in range(min(c1[0], c2[0]), max(c1[0], c2[0]) + 1):
        rockCoords.add((x, c1[1]))
      for y in range(min(c1[1], c2[1]), max(c1[1], c2[1]) + 1):
        rockCoords.add((c1[0], y))
  return rockCoords

def findLowestRock(rockCoords):
  max = 0
  for _,y in rockCoords:
    if y > max: max = y
  return max

def pourSand(starting, rockCoords, lowestRock):
  landed = 0 
  sandPositions = set()
  while True:
    sandPosition = (starting, 0)
    while True:
      nextPosition = getNextSandPosition(sandPosition, rockCoords, sandPositions, lowestRock)
      if nextPosition == sandPosition: 
        sandPositions.add(sandPosition)
        landed += 1
        if sandPosition == (starting, 0):
          return landed
        break
      # if nextPosition[1] > lowestRock:
      #   return landed
      sandPosition = nextPosition


def getNextSandPosition(sandPosition, rockCoords, sandCoords, lowestRock):
  movements = getAdjPositions(sandPosition)
  for position in movements:
    if not (position in rockCoords or position in sandCoords or position[1] >= lowestRock + 2 ):
      return position
  return sandPosition



def getAdjPositions(sandPosition):
  movements = [(0, 1), (-1, 1), (1, 1)]
  positions = []
  for dx, dy in movements:
    positions.append((sandPosition[0] + dx, sandPosition[1] + dy))
  return positions

def solve(fileName):
  coords = parseInput(fileName)
  rockCoords = createFlatArrOfCoords(coords)
  lowestRock = findLowestRock(rockCoords)
  sands = pourSand(500, rockCoords, lowestRock)
  print(sands)


  

if __name__ == "__main__":
  main()
