#!/usr/bin/env python3
# Advent of Code Challenge 15

class Sensor:
  def __init__(self, coord, beacon_coord):
    self.coord = coord
    self.beacon_coord = beacon_coord
  
  @property
  def m_distance_to_beacon(self):
    return abs(self.coord[0] - self.beacon_coord[0]) + abs(abs(self.coord[1] - self.beacon_coord[1]))


def getMDistance(c1, c2):
  return abs(c1[0] - c2[0]) + abs(c1[1] - c2[1])

def parseInput(inputFilename):
  sensors = {}
  with open(inputFilename) as f:
    for line in f:
      tokens = line.strip().split(" ")
      sensor = (int(tokens[2][2:-1]), int(tokens[3][2:-1]))
      beacon = (int(tokens[8][2:-1]), int(tokens[9][2:]))
      sensors[sensor] = beacon
  
  return sensors
  for k,v in sensors.items():
    print(k, v)


def solve(inputFilename, rowToCheck):
  sensor_map = parseInput(inputFilename)


def main():
  solve("input.txt", 10)

if __name__ == "__main__":
  main()
