#!/usr/bin/env python3

# Rock     A -> X
# Paper    B -> Y
# Scissors C -> Z

points = {
  "X": 1,
  "Y": 2,
  "Z": 3,
  "A": 1,
  "B": 2,
  "C": 3,
}


def part1(): 
  total = 0
  with open("input.txt", "r") as f:
    for line in f:
      strats = line.strip().split(" ")
      enemy = strats[0]
      play = strats[1]
      sum = points[enemy] + points[play]
      # Draw
      if (points[enemy] == points[play]):
        total += 3
      # Win
      elif ( (sum == 3 and play == "Y") or (sum == 4 and play == "X") or (sum == 5 and play == "Z")):
        total += 6
      total += points[play]
  return total

def part2():
  total = 0
  with open("input.txt", "r") as f:
    for line in f:
      strats = line.strip().split(" ")
      play = strats[0]
      result = strats[1]

      points = getPoint(play, result)
      total += points
  return total

def getPoint(play, result):
  # Lose
  if (result == "X"):
    return ((points[play] - 2) % 3) + 1
  # Draw
  elif (result == "Y"):
    return points[play] + 3
  # Win
  elif (result == "Z"):
    return ( (points[play] + 3) % 3) + 1 + 6
  return 0

def main():
  print("Part 1:", part1())
  print("Part 2:", part2())

  

if __name__ == "__main__":
  main()