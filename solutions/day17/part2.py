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

class RockTower:
  def __init__(self, chamber_width, wind_directions):
    self.chamber_width = chamber_width  
    self.wind_directions = wind_directions
    self.wind_index = -1
    self.rocks_dropped = 0
    self.patterns = {}
    self.highest_position = 0
    self.rock_placements = set()

  def getNextRockCoords(self):
    shape = rock_shapes[self.rocks_dropped % len(rock_shapes)]
    return list(map(lambda x: (x[0], x[1] + self.highest_position + 3 + 1), shape))
  
  def moveRockHorizontally(self, rock):
    new_position = []
    for x, y in rock:
      nx = x + self.wind_directions[self.wind_index]
      if nx < 1 or nx > self.chamber_width or (nx, y) in self.rock_placements:
        return rock
      new_position.append((nx, y))
    return new_position
  
  def moveRockVertically(self, rock):
    new_position = []
    for x, y in rock:
      ny = y - 1
      if ny == 0 or (x, ny) in self.rock_placements:
        return rock, True
      new_position.append((x, ny))
    return new_position, False
  

  def printTower(self, lines_to_print=int(10000000000)):
    print()
    for y in range(self.highest_position, -1, -1):
      line = "|"
      if y == 0:
        line = "|" + ("+" * self.chamber_width) + "|"
        print(line)
        continue
      for x in range (self.chamber_width):
        if (x + 1, y) in self.rock_placements:
          line += "#"
        else:
          line += "·"
      line += "|"
      print(line)
      if self.highest_position - y > lines_to_print:
        break
    print()
    print()
  
  def dropRocks(self, num_rocks: int):
    while self.rocks_dropped < num_rocks:
      rock = self.getNextRockCoords()
      self.dropRock(rock)
      self.rocks_dropped += 1
  
  def distanceToNextRock(self):
    dist = []
    for x in range(1, self.chamber_width + 1):
      for y in range(self.highest_position, -1, -1):
        if (x, y) in self.rock_placements or y == 0 :
          dist.append(self.highest_position - y)
          break
    return dist


  def dropUntilPatternFound(self, max_rocks):
    while True:
      rock = self.getNextRockCoords()
      self.dropRock(rock)
      self.rocks_dropped += 1
      if self.rocks_dropped % 10000 == 0:
        print(self.rocks_dropped)

      key = str(self.rocks_dropped % 5) + "-" + str(self.wind_index) + "-" + ''.join(str(x) for x in self.distanceToNextRock())
      if key in self.patterns.keys():
        print("Initially found cycle with rocks fallen:", self.patterns[key][0])
        print("Initially found cycle at height:", self.patterns[key][1])
        print("Current Rocks Fallen:", self.rocks_dropped)
        print("Current Height:", self.highest_position)
        
        cycle_height = self.highest_position - self.patterns[key][1]
        cycle_rocks = (self.rocks_dropped  - self.patterns[key][0])
        num_cycles = (max_rocks - self.patterns[key][0]) // cycle_rocks
        
        ans_rocks = self.patterns[key][0] + (num_cycles * cycle_rocks)
        remaining_rocks = max_rocks - ans_rocks

        print("Remaining rocks:", remaining_rocks)
        self.rocks_dropped = ans_rocks
        self.dropRocks(max_rocks)

        final_answer = self.patterns[key][1] + (num_cycles * cycle_height) + (self.highest_position - self.patterns[key][1] - cycle_height)
        print("Part 2:", final_answer)
        return
      
      self.patterns[key] = (self.rocks_dropped, self.highest_position)

  def dropRock(self, rock):
    placed = False
    while not placed:
      self.wind_index = (self.wind_index + 1) % len(self.wind_directions)
      rock = self.moveRockHorizontally(rock)
      [rock, placed] = self.moveRockVertically(rock)
    self.rock_placements = self.rock_placements.union(rock)
    nhp = max(map(lambda x: x[1], rock))
    if nhp > self.highest_position: self.highest_position = nhp


def solve(filename, rocks_to_fall, chamber_width, part):
  wind = parseInput(filename)
  tower = RockTower(chamber_width, wind)
  if part == 1:
    tower.dropRocks(rocks_to_fall)
    print("Part 1:", tower.highest_position)
    print()
  elif part == 2:
    tower.dropUntilPatternFound(rocks_to_fall)
  

def main():
  solve("input.txt", 2022, 7, 1)
  solve("input.txt", 1000000000000, 7, 2)

if __name__ == "__main__":
  main()
