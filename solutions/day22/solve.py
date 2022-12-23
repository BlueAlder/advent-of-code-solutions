#!/usr/bin/env python3
# Advent of Code Challenge 22

import os

open_tiles = set()
wall_tiles = set()

numbers = "0123456789"

class Monke:
  directions = [(1, 0), (0, -1), (-1, 0), (1, 0)]
  def __init__(self, position) -> None:
    self.position = position
    self.direction = 0

  def rotate(self, direction):
    if direction == "R":
      self.direction = (self.direction + 1) % len(direction) 
    elif direction == "L":
      self.direction = (self.direction - 1) % len(direction)
    else:
      raise "Bruh"
  
  def move(self) -> bool:
    npos = (self.position[0] + self.directions[self.direction][0], self.position[1] + self.directions[self.direction][1])
    if npos in wall_tiles:
      return False
    if npos not in open_tiles:
      

    

  


def parseInput(filename):
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    y = 1
    while True:
      line = f.readline().rstrip()
      if line == "": break
      for x, chr in enumerate(line):
        if chr == ".":
          open_tiles.add((x + 1, y))
        elif chr == "#":
          wall_tiles.add((x + 1, y))
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

def startWalking(instructions):
  x = min(map(lambda p: p[0], filter(lambda p: p[1] == 1, open_tiles)))
  start_position = (x, 1)
  m = Monke(start_position)

  for instruction in instructions:
    if isinstance(instruction, int):
      pass
    else:
      m.rotate(instruction)

def solve(filename):
  instructions = parseInput(filename)
  print(instructions)
  startWalking(instructions)

def main():
  solve("input.txt")

if __name__ == "__main__":
  main()
