#!/usr/bin/env python3
# Advent of Code Challenge 16

import os
from collections import defaultdict
import copy

class Node:
  def __init__(self, id, value, connecting_nodes):
    self.id = id
    self.value = value
    self.connecting_nodes = connecting_nodes
  

def parseInput(filename):
  nodes = {}
  with open(os.path.join(os.path.dirname(__file__), filename)) as f:
    for line in f:
      toks = line.strip().split(" ")
      valve_id = toks[1]
      rate = int(toks[4][5:-1])
      adj = list(map(lambda x: x.replace(',', ''), toks[9:]))
      nodes[valve_id] = Node(valve_id, rate, adj)
  return nodes

def getConnectingValves(g, node):
  next = []
  for conn in g[node].connecting_nodes:
    next.append(g[conn].id)
  return next



# def findNextDestination(g, starting_node_id, minutes, opened=None):
#   visited = set()
#   queue = []
#   node_steps = defaultdict(lambda: 1e7)
#   node_steps[starting_node_id] = 0
#   queue.append(starting_node_id)

#   while queue:
#     curr_position = queue.pop(0)
#     visited.add(curr_position)

#     for nextPos in getConnectingValves(g, curr_position.id):
#       if nextPos in visited: continue
#       new_steps = node_steps[curr_position] + 1
#       if new_steps < node_steps[nextPos]:
#         node_steps[nextPos] = new_steps
#         queue.append(nextPos)
#   values = {}
#   max_val = 0
#   destination = ""
#   count = 0
#   for node, step in node_steps.items():
#     val = node.value * (minutes - step - 1)
#     if val > 0:
#       count += 1
#     if val > max_val: 
#       max_val = val
#       destination = node
#     print(node.id, step ,val)
#   print(max_val)
#   print(destination.id)
#   print(count)

def findMaxGasReleased(starting_point_id, 
                      distance_between_valves, 
                      valves, 
                      opened_v, 
                      # pressure_released,
                      minutes_remaining): 
  #   return 0
  max_gas_released = 0
  for valve in valves.values():
    # if starting_point_id == "AA":
      # print('s')
    if valve.id in opened_v: continue
    # 1 Is it possible to go there in the remaining amount of time?

    # 2 Can we turn it on and still have time remaining so it will actually do something?
    d = distance_between_valves[starting_point_id + valve.id]
    if d > minutes_remaining + 1: continue
    pressure = valve.value * (minutes_remaining - d - 1) #+ pressure_released
    new_minutes = minutes_remaining - (d + 1)
    new_opened = copy.copy(opened_v)
    new_opened.append(valve.id)
    gas_released = findMaxGasReleased(valve.id, distance_between_valves, valves, new_opened, new_minutes)
    if gas_released + pressure > max_gas_released: max_gas_released = gas_released + pressure
  
  # print("gday")
  # if len(opened_v) == 6 and starting_point_id == "BB":
      # print('s')
  return max_gas_released






def calcDistanceBetweenValvesAndStartPoint(nodes, valves, start_point):
  distances = {}
  for id, valve in valves.items():
    starting_node_id = id
    visited = set()
    queue = []
    node_steps = defaultdict(lambda: 1e7)
    node_steps[starting_node_id] = 0
    queue.append(starting_node_id)

    while queue:
      curr_position = queue.pop(0)
      visited.add(curr_position)
      for nextPos in getConnectingValves(nodes, curr_position):
        if nextPos in visited: continue
        new_steps = node_steps[curr_position] + 1
        if new_steps < node_steps[nextPos]:
          node_steps[nextPos] = new_steps
          queue.append(nextPos)
    
    for key in valves.keys():
      distances[valve.id + key] = node_steps[key]
    distances[valve.id + start_point] = node_steps[start_point] 
    distances[start_point + valve.id] = node_steps[start_point] 
  return distances


def solve(filename, starting_node_id, minutes):
  nodes = parseInput(filename)
  valves = dict(filter(lambda  node: node[1].value > 0, nodes.items()))
  d_between_valves = calcDistanceBetweenValvesAndStartPoint(nodes, valves, starting_node_id)
  m = findMaxGasReleased(starting_node_id, 
                      d_between_valves, 
                      valves, 
                      [], 
                      minutes)
  print("checking")

  
  print(m)

def main():
  solve("input.txt", "AA", 30)

if __name__ == "__main__":
  main()
