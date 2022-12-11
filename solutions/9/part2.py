#!/usr/bin/env python3

import math

def main():
  positions = []
  chain_length = 10
  for i in range(chain_length):
    positions.append([0,0])

  locations = set()
  locations.add(''.join(map(str,[0, 0])))

  with open("input.txt", "r") as f:
    for line in f:
      tokens = line.strip().split(" ")
      direction = tokens[0]
      size = int(tokens[1])
      
      for i in range(size):
        positions[0] = move(direction, positions[0])
        
        for i in range(1, chain_length):
          dis = distance(positions[i - 1], positions[i])
          if dis > 2:
            # positions[i] = move(direction, positions[i])
            if positions[i][0] < positions[i-1][0] and positions[i][1] < positions[i-1][1]:
              positions[i] = move("R", move("U", positions[i]))
            elif positions[i][0] > positions[i-1][0] and positions[i][1] < positions[i-1][1]:
              positions[i] = move("L", move("U", positions[i]))
            elif positions[i][0] > positions[i-1][0] and positions[i][1] > positions[i-1][1]:
              positions[i] = move("L", move("D", positions[i]))
            elif positions[i][0] < positions[i-1][0] and positions[i][1] > positions[i-1][1]:
              positions[i] = move("R", move("D", positions[i]))
            
          elif dis == 2:
            if positions[i][0] > positions[i-1][0]:
              positions[i] = move("L", positions[i])
            elif positions[i][0] < positions[i-1][0]:
              positions[i] = move("R", positions[i])
            elif positions[i][1] > positions[i-1][1]:
              positions[i] = move("D", positions[i])
            else:
              positions[i] = move("U", positions[i])
      
        locations.add(''.join(map(str,positions[-1])))
        if direction == "R" and size == 17:
          printPositions(positions)
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

def printPositions(positions):
  print()
  print("==")
  for y in range(10, -10, -1):
    xStr = ""
    for x in range(-10, 20):
      if [x, y] in positions:
        xStr += str(positions.index([x, y])) + " "
      else:
        xStr += ". "
    print(xStr)

if __name__ == "__main__":
  main()