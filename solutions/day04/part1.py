#!/usr/bin/env python3

def main():
  count = 0
  with open("input.txt", "r") as f:
    for pair in f:
      sections = pair.strip().split(",")

      pair1lower = int(sections[0].split("-")[0])
      pair1upper = int(sections[0].split("-")[1])

      pair2lower = int(sections[1].split("-")[0])
      pair2upper = int(sections[1].split("-")[1])

      if pair1lower <= pair2lower and pair1upper >= pair2upper:
        count += 1
      elif pair1lower >= pair2lower and pair1upper <= pair2upper:
        count += 1
  print(count)
if __name__ == "__main__":
  main()