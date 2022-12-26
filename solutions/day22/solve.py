#!/usr/bin/env python3
# Advent of Code Challenge 22

import os

tiles = set()
numbers = "0123456789"

def findWrap(ordinate: int , horizontal: bool, wrap_to_start: bool):
  line = map(lambda p: p[0 if horizontal else 1], filter(lambda p: p[1 if horizontal else 0] == ordinate, tiles))
  val = min(line) if wrap_to_start else max(line)
  return (val, ordinate) if horizontal else (ordinate, val)

class Monke:
  directions = [(1, 0), (0, 1), (-1, 0), (0, -1)]
  def __init__(self, position) -> None:
    self.position = position
    self.direction = 0
  
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
          if not res: break
      else:
        self.rotate(instruction)
  
  def move(self) -> bool:
    npos = (self.position[0] + self.directions[self.direction][0], self.position[1] + self.directions[self.direction][1], True)
    if npos in tiles:
      self.position = npos
      return True
    else:
      npos = (npos[0], npos[1], False)
      if npos in tiles: return False

    # Wrap
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


def solve(filename):
  instructions = parseInput(filename)
  x = min(map(lambda p: p[0], filter(lambda p: p[1] == 1 and p[2] == True, tiles)))
  start_position = (x, 1, True)
  m = Monke(start_position)
  m.startWalking(instructions)
  print(m)
  sol = (m.position[1] * 1000) + (m.position[0] * 4) + m.direction
  print(sol)

def main():
  solve("input.txt")

if __name__ == "__main__":
  main()
