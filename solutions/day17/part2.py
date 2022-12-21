#!/usr/bin/env python3
# Advent of Code Challenge 17
import os
from tqdm import tqdm


# 1 Build until we get a perfect floor
# 2 Memoise the wind index, rock_index and any above floor rocks
# 3 Continue building and memoising until we hit a duplicate (wind_index, rock_index, above floor rocks)

def parseInput(filename):
  wind = []
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    data = f.readline().strip()
    for w in data:
      wind.append(-1) if w == "<" else wind.append(1)
  return wind

rock_shapes = [
  [(3, 0), (4, 0), (5, 0), (6, 0)],
  [(3, 1), (4, 0), (4, 1), (4, 2), (5,1)],
  [(3, 0), (4, 0), (5, 0), (5, 1), (5,2)],
  [(3, 0), (3, 1), (3, 2), (3, 3)],
  [(3, 0), (3, 1), (4, 0), (4, 1)],
]

def getNextRockCoords(rock_index, highest_position):
  shape = rock_shapes[rock_index % len(rock_shapes)]
  return list(map(lambda x: (x[0], x[1] + highest_position + 3 + 1), shape))

def pushRock(direction, rock, rocks, chamber_width):
  new_position = []
  for x, y in rock:
    nx = x + direction
    if nx < 1 or nx > chamber_width or (nx, y) in rocks:
      return rock
    new_position.append((nx, y))
  return new_position

def dropRock(rock, rocks):
  new_position = []
  for x, y in rock:
    ny = y - 1
    if ny == 0 or (x, ny) in rocks:
      return rock, True
    new_position.append((x, ny))
  return new_position, False
  
def printTower(width, max_height, rocks):
  print()
  for y in range(max_height, -1, -1):
    line = "|"
    if y == 0:
      line = "|" + ("+" * width) + "|"
      print(line)
      continue
    for x in range (width):
      if (x + 1, y) in rocks:
        line += "#"
      else:
        line += "."
    line += "|"
    print(line)
  print()
  print()



def solve(filename, rocks_to_fall, chamber_width):
  wind = parseInput(filename)
  current_wind_index = 0
  highest_position = 0
  # Floor will be at y = 0
  # wall will be at x = 0 and x = 8
  rock_coords = set()
  for rock_index in tqdm(range(rocks_to_fall)):
    rock = getNextRockCoords(rock_index, highest_position)
    # print("Rock", str(rock_index + 1) + ":", rock)
    placed = False
    while not placed:
      rock = pushRock(wind[current_wind_index], rock, rock_coords, chamber_width)
      # print("Rock", str(rock_index + 1) + ":", rock)
      [rock, placed] = dropRock(rock, rock_coords)
      # print("Rock", str(rock_index + 1) + ":", rock)

      # might be bug with wind index
      current_wind_index = (current_wind_index + 1) % len(wind)
      if placed:
        rock_coords = rock_coords.union(rock)
        nhp = max(map(lambda x: x[1], rock))
        if nhp > highest_position: highest_position = nhp
        # printTower(chamber_width, highest_position, rock_coords)

  print(highest_position)

    # break

    
  


def main():
  solve("input.txt", 1000000000000, 7)

if __name__ == "__main__":
  main()
