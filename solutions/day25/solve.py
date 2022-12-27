#!/usr/bin/env python3
# Advent of Code Challenge 25

import os

digits = ["=", "-", "0", "1", "2"]

def parseInput(filename):
  snafus = []
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for line in f:
      snafus.append(line.strip())
  return snafus
  
def snafuToInt(snafu: str) -> int:
  decimal = 0
  for i, chr in enumerate(snafu):
    decimal += (digits.index(chr) - 2) * (5 ** (len(snafu) - i - 1))
  return decimal

def intToSnafu(decimal: int) -> str:
  if decimal == 8:
    print("hih")
  res = ""
  next_offset = 2
  current =  decimal

  idx = 0 
  while current > 0:
    q, mod  = divmod(current, 5)
    # q, m = divmod(mod + 2, 5)\
    i = (mod + 2) % 5
    res = digits[i] + res
    if i < 2:
      current += (2 - i) * (5 ** idx)   #- 2
    else:
      current = q
    idx += 1
  return res






def solve(filename):
  snafus = parseInput(filename)
  total = 0
  for snafu in snafus:
    total += snafuToInt(snafu)
  print(total)
  
  for i in range(1, 10):
    print(intToSnafu(i))


def main():
  solve("input.txt")

if __name__ == "__main__":
  main()
