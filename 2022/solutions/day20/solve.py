#!/usr/bin/env python3
# Advent of Code Challenge 20

import os

class Value:
  def __init__(self, value):
    self.value = value
    self.prev: Value = None
    self.next: Value = None
  
def parseInput(filename, dec_key):
  arr = []
  zero_val = None
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for dig in f:
      v = Value(int(dig.strip()) * dec_key)
      arr.append(v)
      if v.value == 0: zero_val = v
  for i, v in enumerate(arr):
    v.prev = arr[(i - 1) % len(arr)]
    v.next = arr[(i + 1) % len(arr)]
  return arr, zero_val

def swap(p: Value, q: Value):
  q.prev = p.prev
  p.prev.next = q
  
  q.next.prev = p
  p.next = q.next

  p.prev = q
  q.next = p

# I definitely could optimise this by not swapping every element and rather directly
# inserting, but oh well. It works  
def unmix(mixed):
  for val in mixed:
    iter =  abs(val.value) % (len(mixed) - 1) 
    for i in range(iter):
      swap(val, val.next) if val.value > 0 else swap(val.prev, val)

def solve(filename, dec_key, mixes):
  mixed, zero_val = parseInput(filename, dec_key)
  for i in range(mixes):
    unmix(mixed)
    print(i + 1, "mix done")

  val = zero_val
  sum = 0
  for i in range(3):
    for i in range(1000):
      val = val.next
    sum += val.value
  return sum

def main():
  p1 = solve("input.txt", 1, 1)
  print("Part 1:", p1)
  p2 = solve("input.txt", 811589153, 10)
  print("Part 2:", p2)

if __name__ == "__main__":
  main()
