#!/usr/bin/env python3

import math

def main():
  head = [0, 0]
  tail = [0, 0]

  locations = set()
  locations.add(''.join(map(str,tail)))

  with open("input.txt", "r") as f:
    for line in f:
      tokens = line.strip().split(" ")
      direction = tokens[0]
      size = int(tokens[1])
      for i in range(size):
        head = move(direction, head)
        dis = distance(head, tail)
        if dis > 2:
          tail = move(direction, tail)
          if direction == "R" or direction == "L":
            if tail[1] > head[1]:
              tail = move("D", tail)
            else:
              tail = move("U", tail)
            locations.add(''.join(map(str,tail)))
          else:
            if tail[0] > head[0]:
              tail = move("L", tail)
            else:
              tail = move("R", tail)
            locations.add(''.join(map(str,tail)))


        elif dis == 2:
          tail = move(direction, tail)
          locations.add(''.join(map(str,tail)))

        # print(head)
    print(len(locations))
        
  pass

def distance(p1, p2):
  xdist = abs(p1[0] - p2[0])
  ydist = abs(p1[1] - p2[1])

  c2 = (xdist** 2) + (ydist ** 2)
  return math.sqrt(c2)

def move(dir, pointer):
  if   dir == "R":
    pointer[0] += 1
  elif dir == "L":
    pointer[0] -= 1
  elif dir == "U":
    pointer[1] += 1
  elif dir == "D":
    pointer[1] -= 1
  else:
    raise "Bad direction"
  return pointer

if __name__ == "__main__":
  main()