#!/usr/bin/env python3

def main():
  total = 0
  with open("input.txt", "r") as f:
    for r1, r2, r3 in zip(f, f, f):
      letter = findCommonLetter(r1.strip(), r2.strip(), r3.strip())
      total += getPriorityFromLetter(letter)
  print(total)


def findCommonLetter(r1, r2, r3):
  for chr in r1:
    if chr in r2 and chr in r3:
      return chr

def getPriorityFromLetter(letter):
  if letter.isupper():
    return ord(letter) - 64 + 26
  else:
    return ord(letter) - 96

if __name__ == "__main__":
  main()