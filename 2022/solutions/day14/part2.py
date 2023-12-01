#!/usr/bin/env python3
# Advent of Code Challenge 14

import itertools
import os
import time
import sys

sys.setrecursionlimit(10000)


def parseInput(fileName):
  coords = []
  with open(os.path.join(os.path.dirname(__file__), fileName)) as f:
    for line in f:
      rockLine = line.strip().split(" -> ")
      rockLine = list(map(lambda x: (int(x.split(",")[0]), int(x.split(",")[1])), rockLine))
      coords.append(rockLine)
  return coords
  
def createFlatArrOfCoords(coords):
  rockCoords = set()
  for coord in coords:
    for c1, c2 in list(zip(coord, coord[1:])):
      for x in range(min(c1[0], c2[0]), max(c1[0], c2[0]) + 1):
        rockCoords.add((x, c1[1]))
      for y in range(min(c1[1], c2[1]), max(c1[1], c2[1]) + 1):
        rockCoords.add((c1[0], y))
  return rockCoords

# BFS/DFS to find all possible positions for sand to go.
def pourSand(starting, rockCoords, floor, dfs=False):
  visited = set()
  queue = []
  queue.append(starting)
  visited.add(starting)

  while queue:
    sandCoord = queue.pop() if dfs else queue.pop(0)

    for nextPos in getPossibleNextPositions(sandCoord, rockCoords, floor):
      if nextPos in visited: continue
      queue.append(nextPos)
      visited.add(nextPos)
  return len(visited)


def getPossibleNextPositions(sandPosition, rockCoords, floor):
  movements = getAdjPositions(sandPosition)
  positions = []
  for position in movements:
    if not (position in rockCoords or position[1] >= floor):
      positions.append(position)
  return positions


def getAdjPositions(sandPosition):
  movements = [(0, 1), (-1, 1), (1, 1)]
  positions = []
  for dx, dy in movements:
    positions.append((sandPosition[0] + dx, sandPosition[1] + dy))
  return positions

def solve(fileName, startCoord, floor_delta):
  coords = parseInput(fileName)
  rockCoords = createFlatArrOfCoords(coords)
  floor = max(map(lambda x: x[1], rockCoords)) + floor_delta
  
  start_time = time.time()
  sandsBFS = pourSand(startCoord, rockCoords, floor)
  print("BFS:", sandsBFS, time.time() - start_time, "seconds")
  
  start_time = time.time()
  sandsDFS = pourSand(startCoord, rockCoords, floor, True)
  print("DFS:", sandsDFS, time.time() - start_time, "seconds")



def main():
  solve("input.txt", (500, 0), 2)

if __name__ == "__main__":
  main()
