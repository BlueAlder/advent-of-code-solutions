#!/usr/bin/env python3
# Advent of Code Challenge 22

import os

# Please hardcoding gods do not hate upon me

EDGES = [
  {
    # A
    "x_range": [51, 100],
    "y_range":  [1, 1],
    "direction": 3,
    "offset": 100,
    "new_direction": 0,
    "new_x": 1,
    "new_y": None,
    "invert": False,
    "checked": False
  },
  {
    # B
    "x_range": [101, 150],
    "y_range":  [1, 1],
    "direction": 3,
    "offset": -100,
    "new_direction": 3,
    "new_x": None,
    "new_y": 200,
    "invert": False,
    "checked": False
    
  },
  {
    # C
    "x_range": [150, 150],
    "y_range":  [1, 50],
    "direction": 0,
    "offset": 100,
    "new_direction": 2,
    "new_x": 100,
    "new_y": None,
    "invert": True,
    "checked": False
  },
  {
    # D
    "x_range": [101, 150],
    "y_range":  [50, 50],
    "direction": 1,
    "offset": -50,
    "new_direction": 2,
    "new_x": 100,
    "new_y": None,
    "invert": False,
    "checked": False
  },
  {
    # E
    "x_range": [100, 100],
    "y_range":  [51, 100],
    "direction": 0,
    "offset": 50,
    "new_direction": 3,
    "new_x": None,
    "new_y": 50,
    "invert": False,
    "checked": False
  },
  {
    # F
    "x_range": [100, 100],
    "y_range":  [101, 150],
    "direction": 0,
    "offset": 100,
    "new_direction": 2,
    "new_x": 150,
    "new_y": None,
    "invert": True,
    "checked": False
  },
  {
    # G
    "x_range": [51, 100],
    "y_range":  [150, 150],
    "direction": 1,
    "offset": 100,
    "new_direction": 2,
    "new_x": 50,
    "new_y": None,
    "invert": False,
    "checked": False
  },
  {
    # H
    "x_range": [50, 50],
    "y_range":  [151, 200],
    "direction": 0,
    "offset": -100,
    "new_direction": 3,
    "new_x": None,
    "new_y": 150,
    "invert": False,
    "checked": False
  },
  {
    # I
    "x_range": [1, 50],
    "y_range":  [200, 200],
    "direction": 1,
    "offset": 100,
    "new_direction": 1,
    "new_x": None,
    "new_y": 1,
    "invert": False,
    "checked": False
  },
  {
    # J
    "x_range": [1, 1],
    "y_range":  [151, 200],
    "direction": 2,
    "offset": -100,
    "new_direction": 1,
    "new_x": None,
    "new_y": 1,
    "invert": False,
    "checked": False
  },
  {
    # K
    "x_range": [1, 1],
    "y_range":  [101, 150],
    "direction": 2,
    "offset": 100,
    "new_direction": 0,
    "new_x": 51,
    "new_y": None,
    "invert": True,
    "checked": False
  },
  {
    # L
    "x_range": [1, 50],
    "y_range":  [101, 101],
    "direction": 3,
    "offset": 50,
    "new_direction": 0,
    "new_x": 51,
    "new_y": None,
    "invert": False,
    "checked": False
  },
  {
    # M
    "x_range": [51, 51],
    "y_range":  [51, 100],
    "direction": 2,
    "offset": -50,
    "new_direction": 1,
    "new_x": None,
    "new_y": 101,
    "invert": False,
    "checked": False
  },
  {
    # N
    "x_range": [51, 51],
    "y_range":  [1, 50],
    "direction": 2,
    "offset": 100,
    "new_direction": 0,
    "new_x": 1,
    "new_y": None,
    "invert": True,
    "checked": False
  }
]


tiles = set()
numbers = "0123456789"

def findWrap(ordinate: int , horizontal: bool, wrap_to_start: bool):
  line = map(lambda p: p[0 if horizontal else 1], filter(lambda p: p[1 if horizontal else 0] == ordinate, tiles))
  val = min(line) if wrap_to_start else max(line)
  return (val, ordinate) if horizontal else (ordinate, val)

class Monke:
  directions = [(1, 0), (0, 1), (-1, 0), (0, -1)]
  def __init__(self, position, part) -> None:
    self.position = position
    self.direction = 0
    self.part = part
  
  def __str__(self) -> str:
    return f"{self.position} {self.direction}"

  def rotate(self, direction):
    if direction == "R":
      self.direction = (self.direction + 1) % len(self.directions) 
    elif direction == "L":
      self.direction = (self.direction - 1) % len(self.directions)
    else:
      raise "Bruh"
  
  def startWalking(self, instructions):
    for instruction in instructions:
      print(f"Instruction: {instruction}")
      if isinstance(instruction, int):
        for i in range(instruction):
          res = self.move()
          # print(self)
          if not res: break
      else:
        self.rotate(instruction)
      if (self.position == (88, 54, True)):
        print("hi")
      print(self)
  
  def move(self) -> bool:
    npos = (self.position[0] + self.directions[self.direction][0], self.position[1] + self.directions[self.direction][1], True)
    if npos in tiles:
      self.position = npos
      return True
    else:
      npos = (npos[0], npos[1], False)
      if npos in tiles: return False

    # Wrap

    # 1. Find which edge we are on
    if self.part == 2:
      edge = None
      for pedge in EDGES:
        if (self.position[0] in range(pedge["x_range"][0], pedge["x_range"][1] + 1) and 
            self.position[1] in range(pedge["y_range"][0], pedge["y_range"][1] + 1) and
            self.direction == pedge["direction"]):
            edge = pedge
            break
      # print(edge)
      
      swapxy = (edge["direction"] + edge["new_direction"]) % 2 != 0 and edge["direction"] != edge["new_direction"]
      nx = self.position[0] if not swapxy else self.position[1]
      ny = self.position[1] if not swapxy else self.position[0]
      if edge["new_x"]: 
        nx = edge["new_x"]
        ny = ny + edge["offset"] if not edge["invert"] else (50 - ny + 1) + edge["offset"]
      if edge["new_y"]: 
        ny = edge["new_y"]
        nx = nx + edge["offset"] if not edge["invert"] else (50 - nx + 1) + edge["offset"]
      
      if not edge["checked"]:
        edge["checked"] = True

      # print(nx, ny)
      npos = (nx, ny, True)
      print((nx, ny))
      if npos in tiles:
        self.position = npos
        self.direction = edge["new_direction"]
        return True
      return False

    if self.part == 1 :
      row = True if self.direction % 2 == 0 else False
      ord = self.position[1] if row else self.position[0]
      wrap_to_start = True if self.direction < 2 else False
      w = findWrap(ord, row, wrap_to_start)
      npos = (w[0], w[1], True)
      if npos in tiles:
        self.position = npos
        return True
      return False      

def parseInput(filename):
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    y = 1
    while True:
      line = f.readline().rstrip()
      if line == "": break
      for x, chr in enumerate(line):
        if chr == ".":
          tiles.add((x + 1, y, True))
        elif chr == "#":
          tiles.add((x + 1, y, False))
      y += 1
    instructionsStr = f.readline()
    instructions = []
    curr_tok = ""
    for chr in instructionsStr:
      if chr in numbers:
        curr_tok += chr
      else:
        instructions.append(int(curr_tok))
        instructions.append(chr)
        curr_tok = ""
    if curr_tok: instructions.append(int(curr_tok))
    return instructions


def solve(filename, part):
  instructions = parseInput(filename)
  x = min(map(lambda p: p[0], filter(lambda p: p[1] == 1 and p[2] == True, tiles)))
  start_position = (x, 1, True)
  m = Monke(start_position, part)
  m.startWalking(instructions)
  return (m.position[1] * 1000) + (m.position[0] * 4) + m.direction

def main():
  # p1 = solve("input.txt", 1)
  # print("Part 1:", p1)
  p2 = solve("input.txt", 2)
  print("Part 2:", p2)


if __name__ == "__main__":
  main()
