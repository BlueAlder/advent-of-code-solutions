#!/usr/bin/env python3

def main():
  p1 = 0
  p2 = 0
  with open("input.txt", "r") as f:
    for pair in f:
      sections = pair.strip().split(",")

      pair1lower = int(sections[0].split("-")[0])
      pair1upper = int(sections[0].split("-")[1])

      pair2lower = int(sections[1].split("-")[0])
      pair2upper = int(sections[1].split("-")[1])

      if pair1lower <= pair2lower and pair1upper >= pair2upper:
        p1 += 1
      elif pair1lower >= pair2lower and pair1upper <= pair2upper:
        p1 += 1

      if pair1lower <= pair2upper and pair1upper >= pair2lower:
        p2 += 1

  print("Part 1:", p1)
  print("Part 2:", p2)


if __name__ == "__main__":
  main()