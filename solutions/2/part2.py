#!/usr/bin/env python3

# Rock     A
# Paper    B
# Scissors C

# Lose     X
# Draw     Y
# Win      Z

points = {
  "A": 1,
  "B": 2,
  "C": 3,
}

def main():
  total = 0
  with open("input.txt", "r") as f:
    for line in f:
      strats = line.strip().split(" ")
      play = strats[0]
      result = strats[1]

      points = getPoint(play, result)
      total += points
  print(total)

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



if __name__ == "__main__":
  main()