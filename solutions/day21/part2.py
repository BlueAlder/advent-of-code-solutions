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

def doInvertOperation(a, b, op):
  if op == "+":
    return a - b
  elif op == "-":
    return a + b
  elif op == "*":
    return int(a / b)
  elif op == "/":
    return a * b
  raise "bruh moment"

def findValue(monke, known, unknown):
  if monke == "humn": return 
  if monke in known.keys(): return known[monke]
  m1 = findValue(unknown[monke][0], known, unknown)
  m2 = findValue(unknown[monke][2], known, unknown)
  if m1 != None: known[unknown[monke][0]] = m1
  if m2 != None: known[unknown[monke][2]] = m2
  if None in [m1, m2]: return

  return doOperation(m1, m2, unknown[monke][1])

def solve(filename, monke):
  known, unknown = parseInput(filename)
  m1 = unknown[monke][0]
  m2 = unknown[monke][2]
  s1 = findValue(m1, known, unknown) 
  s2 = findValue(m2, known, unknown)

  m = m1 if s1 == None else m2
  val = s1 if s2 == None else s2
  # v1 = findValue(unknown[monke][0], known, unknown)
  # v2 = findValue(unknown[monke][2], known, unknown)
  # m, val = unknown[monke][0], v1 if v1 != None else unknown[monke][2], v2

  while True:
    m1 = unknown[m][0]
    m2 = unknown[m][2]
    s1 = findValue(m1, known, unknown) 
    s2 = findValue(m2, known, unknown)

    if s1 == None: s1 = val
    if s2 == None: s2 = val

    
    val = doInvertOperation(s1, s2, unknown[m][1])
    if m2 == "jmws": 
      print("hi")
    m = m1 if s1 == None else m2
    if m == "humn":
      print(val)
      break

    


  # res = findValue(monke, known, unknown)

def main():
  solve("input.txt", "root")

if __name__ == "__main__":
  main()
