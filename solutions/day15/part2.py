#!/usr/bin/env python3
# Advent of Code Challenge 15
from collections import defaultdict
import os


class Sensor:
  def __init__(self, coord, beacon_coord):
    self.coord = coord
    self.beacon_coord = beacon_coord
  
  @property
  def m_distance_to_beacon(self):
    return abs(self.coord[0] - self.beacon_coord[0]) + abs(abs(self.coord[1] - self.beacon_coord[1]))


def getMDistance(c1, c2):
  return abs(c1[0] - c2[0]) + abs(c1[1] - c2[1])

def parseInput(filename):
  sensors = {}
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for line in f:
      tokens = line.strip().split(" ")
      sensor = (int(tokens[2][2:-1]), int(tokens[3][2:-1]))
      beacon = (int(tokens[8][2:-1]), int(tokens[9][2:]))
      sensors[sensor] = beacon
  
  # for k,v in sensors.items():
  #   print(k, v)
  return sensors

# def findBeaconless(sensor_map, row_to_check):
#   cannot_go = set()
#   for sensor, beacon in sensor_map.items():
#     md = getMDistance(sensor, beacon)
#     dist_to_row = abs(sensor[1] - row_to_check)

#     if dist_to_row <= md:
#       delta = abs(dist_to_row - md) 
#       startX = sensor[0] - delta 
#       endX = sensor[0] + delta
#       for x in range(startX, endX + 1):
#         p = (x, row_to_check)
#         if not ( p in sensor_map.values()):
#           cannot_go.add(p)

def walkPerimeterPoints(point, mdistance, max, candidates):
  directions = [(1, 1), (-1, 1), (-1, -1), (1, -1)]
  perim_point = (point[0], point[1] - mdistance - 1)
  for dx, dy in directions:
    for n in range(mdistance + 1):
      perim_point = (perim_point[0] + dx, perim_point[1] + dy)
      if perim_point[0] in range(max) and perim_point[1] in range(max):
        candidates[perim_point].add((dx, dy))
        if len(candidates[perim_point]) == 4:
          print(perim_point)
          print("Signal:", (perim_point[0] * 4000000) + perim_point[1])

      




def findMissingBeacon(sensor_map, max):
  potential_beacon_coords = defaultdict(lambda: set())
  for sensor in sensor_map.keys():
    mdistance = getMDistance(sensor, sensor_map[sensor])
    walkPerimeterPoints(sensor, mdistance,  max, potential_beacon_coords)


  
  # print(cannot_go)

def solve(inputFilename):
  sensor_map = parseInput(inputFilename)
  findMissingBeacon(sensor_map, 4000000)
  


def main():
  solve("input.txt")

if __name__ == "__main__":
  main()
