#!/usr/bin/env python3
# Advent of Code Challenge 21

import os

def parseInput(filename):
  known = {}
  unknown = {}
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for line in f:
      toks = line.strip().split(" ")
      monke = toks[0][:4]
      if len(toks) == 4:
        unknown[monke] = [*toks[1:]]
      elif len(toks) == 2:
        known[monke] = int(toks[1])
  return known, unknown

def doOperation(a, b, op):
  if op == "+":
    return a + b
  elif op == "-":
    return a - b
  elif op == "*":
    return a * b
  elif op == "/":
    return int(a / b)
  raise "bruh moment"

def findValue(monke, known, unknown):
  if monke in known.keys(): return known[monke]
  m1 = findValue(unknown[monke][0], known, unknown)
  known[unknown[monke][0]] = m1
  m2 = findValue(unknown[monke][2], known, unknown)
  known[unknown[monke][2]] = m2
  return doOperation(m1, m2, unknown[monke][1])

def solve(filename, monke):
  known, unknown = parseInput(filename)
  res = findValue(monke, known, unknown)
  print("Part 1:", res)

def main():
  solve("input.txt", "root")

if __name__ == "__main__":
  main()
