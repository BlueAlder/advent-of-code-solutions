#!/usr/bin/env python3
# Advent of Code Challenge 17
import os


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
  current_wind_index = -1
  highest_position = 0
  patterns = {}

  # Floor will be at y = 0
  # wall will be at x = 0 and x = 8
  rock_coords = set()
  for rock_index in range(rocks_to_fall):
    rock = getNextRockCoords(rock_index, highest_position)
    placed = False
    while not placed:
      current_wind_index = (current_wind_index + 1) % len(wind)
      rock = pushRock(wind[current_wind_index], rock, rock_coords, chamber_width)
      [rock, placed] = dropRock(rock, rock_coords)

    # Check for same rock index and wind direction
    rock_coords = rock_coords.union(rock)
    nhp = max(map(lambda x: x[1], rock))
    if nhp > highest_position: highest_position = nhp
    key = str((rock_index) % 5) + "-" + str(current_wind_index)
    if key in patterns.keys():
        # we have found a pattern
        print("Initially found with rocks fallen:", patterns[key][0])
        print("Initially found at height:", patterns[key][1])
        print("Current Rocks Fallen:", rock_index + 1)
        print("Current Height:", highest_position)
        cycle_height = highest_position - patterns[key][1]
        cycle_rocks = (rock_index - patterns[key][0])
        num_cycles = (rocks_to_fall - patterns[key][0]) // cycle_rocks
        print(cycle_height, cycle_rocks, num_cycles)

        ans = patterns[key][1] + (num_cycles * cycle_height) # + remaining 
        ans_rocks = patterns[key][0] + (num_cycles * cycle_rocks)
        print(ans, ans_rocks)


        return
    else:
        patterns[key] = (rock_index + 1, highest_position)

  print(highest_position)

    # break

    
  


def main():
  solve("input.txt", 1000000000000, 7)

if __name__ == "__main__":
  main()
