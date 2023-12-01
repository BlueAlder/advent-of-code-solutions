#!/usr/bin/env python3
# Advent of Code Challenge 18

import os

def parseInput(filename):
  points = set()
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for line in f:
      point = line.strip().split(",")
      points.add((int(point[0]), int(point[1]), int(point[2])))
  return points


def getMinAndMaxZValue(points):
  # All points are positive 
  min = 99999999999999999999999
  max = 0
  for point in points:
    z = point[2]
    if z < min: min = z
    if z > max: max = z
  return min, max

def getAdjacentPoints(point):
  movements = [(1, 0, 0), (-1, 0, 0), (0, 1, 0), (0, -1, 0), (0, 0, 1), (0, 0, -1)]
  adjPoints = []
  for dx, dy, dz in movements:
    adjPoints.append((point[0] + dx, point[1] + dy, point[2] + dz))
  return adjPoints

def checkInAirPocket(point, points, minZ, maxZ):
  stack = [point]
  visted = set()
  while stack:
    p = stack.pop()
    visted.add(p)
    if p[2] > maxZ or p[2] < minZ:
      return False
    aPs = getAdjacentPoints(p)
    for ap in aPs:
      if ap not in points and ap not in visted:
        stack.append(ap)
  return True


def findSurfaceArea(points, part):
  min, max = getMinAndMaxZValue(points)
  surfaceArea = 0
  for z in range(min, max + 1):
    # get points with that z value
    lPoints = list(filter(lambda p: p[2] == z, points))
    for point in lPoints:
      adjPoints = getAdjacentPoints(point)
      for aPoint in adjPoints:
        if aPoint not in points:
          if part == 1 or (part == 2 and not checkInAirPocket(aPoint, points, min, max)): surfaceArea += 1 
            

  return surfaceArea


def solve(filename):
  points = parseInput(filename)
  sa = findSurfaceArea(points, 1)
  print("Part 1:", sa)
  sa = findSurfaceArea(points, 2)
  print("Part 2:", sa)


def main():
  solve("input.txt")

if __name__ == "__main__":
  main()
