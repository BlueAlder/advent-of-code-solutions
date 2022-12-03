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

def main():
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
  print(total)




if __name__ == "__main__":
  main()