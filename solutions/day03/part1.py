#!/usr/bin/env python3

def main():
  total = 0
  with open("input.txt", "r") as f:
    for rucksack in f:
      rucksack = rucksack.strip()
      total += getPriority(rucksack)
  print(total)


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

if __name__ == "__main__":
  main()