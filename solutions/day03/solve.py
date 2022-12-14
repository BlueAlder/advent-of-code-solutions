#!/usr/bin/env python3

def part1():
  total = 0
  with open("input.txt", "r") as f:
    for rucksack in f:
      rucksack = rucksack.strip()
      total += getPriority(rucksack)
  return total 

def part2():
  total = 0
  with open("input.txt", "r") as f:
    for r1, r2, r3 in zip(f, f, f):
      letter = findCommonLetter(r1.strip(), r2.strip(), r3.strip())
      total += getPriorityFromLetter(letter)
  return total

def findCommonLetter(r1, r2, r3):
  for chr in r1:
    if chr in r2 and chr in r3:
      return chr

def getPriority(rucksack):
  compartment1 = rucksack[0:len(rucksack) // 2]
  compartment2 = rucksack[len(rucksack) // 2::]

  for chr in compartment2:
    if chr in compartment1: 
      return getPriorityFromLetter(chr)

def getPriorityFromLetter(letter):
  if letter.isupper():
    return ord(letter) - 64 + 26
  else:
    return ord(letter) - 96

def main():
  print("Part 1:", part1())
  print("Part 2:", part2())

if __name__ == "__main__":
  main()