#!/usr/bin/env python3
# Advent of Code Challenge 21

import os

known = {}
unknown = {}

def parseInput(filename):
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for line in f:
      toks = line.strip().split(" ")
      monke = toks[0][:4]
      if len(toks) == 4:
        unknown[monke] = [*toks[1:]]
      elif len(toks) == 2:
        known[monke] = int(toks[1])

def doOperation(a, b, op, invert=False):
  if op == "+":
    return a + b if invert == False else a - b
  elif op == "-":
    return a - b if invert == False else a + b
  elif op == "*":
    return a * b if invert == False else int(a / b)
  elif op == "/":
    return int(a / b) if invert == False else a * b
  raise "bruh moment"


def findValue(monke):
  if monke == "humn": return 
  if monke in known.keys(): return known[monke]
  m1 = findValue(unknown[monke][0])
  m2 = findValue(unknown[monke][2])
  if m1 != None: known[unknown[monke][0]] = m1
  if m2 != None: known[unknown[monke][2]] = m2
  if None in [m1, m2]: return

  return doOperation(m1, m2, unknown[monke][1])

def solve(filename, monke):
  parseInput(filename)
  
  m1 = unknown[monke][0]
  m2 = unknown[monke][2]
  s1 = findValue(m1) 
  s2 = findValue(m2)

  m = m1 if s1 == None else m2
  desired_value = s1 if s2 == None else s2

  while True:
    m1 = unknown[m][0]
    m2 = unknown[m][2]
    s1 = findValue(m1) 
    s2 = findValue(m2)

    op = unknown[m][1]
    if s1:
      if op in ["-", "/"]:
        desired_value = doOperation(s1, desired_value, op)
      else:
        desired_value =  doOperation(desired_value, s1, op, True)
    else:
      desired_value = doOperation(desired_value, s2, op, True)

    m = m1 if s1 == None else m2
    if m == "humn":
      return desired_value

def main():
  p2 = solve("input.txt", "root")
  print("Part 2:", p2)

if __name__ == "__main__":
  main()
