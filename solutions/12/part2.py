#!/usr/bin/env python3 

from collections import defaultdict
import heapq as heap

class Point:
  def __init__(self, x, y, value):
    self.x = x
    self.y = y
    self.value = value
  
class Game:
  def __init__(self):
    self.grid = []
    with open("input.txt", "r") as f:
      for y, line in enumerate(f):
        row = []
        for x, chr in enumerate(line.strip()):
          if chr == "S" :
            row.append(Point(x, y, 1))
          elif chr == "E":
            self.end = Point(x, y, 26)
            row.append(self.end)
          else:
            row.append(Point(x, y, ord(chr) - 96))
        self.grid.append(row)

  @property
  def width(self):
    return len(self.grid[0])
  
  @property
  def height(self):
    return len(self.grid)


  def availablePoints(self, current_point):
    point_val = current_point.value
    points = []
    movements = [(0,1), (0, -1), (1, 0), (-1, 0)]

    for mx, my in movements:
      test_x = mx + current_point.x
      test_y = my + current_point.y
      if test_x in range(self.width) and test_y in range(self.height) and self.grid[test_y][test_x].value + 1 >= point_val:
        points.append(self.grid[test_y][test_x])
    return points

  def dijkstra(self, start_node, destination_node_value):
    visited = set()
    pq = []
    node_costs = defaultdict(lambda: 1e7)
    node_costs[start_node] = 0
    entry_count = 0
    heap.heappush(pq, (0, entry_count, start_node))

    while pq:
      _, _, node = heap.heappop(pq)
      # Return on first pop of possible start point as it is the shortest.
      if node.value == destination_node_value:
        return node_costs[node]
      visited.add(node)

      for adjPoint in self.availablePoints(node):
        if adjPoint in visited: continue
        entry_count += 1
        new_cost = node_costs[node] + 1
        if node_costs[adjPoint] > new_cost:
          node_costs[adjPoint] = new_cost
          heap.heappush(pq, (new_cost, entry_count, adjPoint))
  
  def bfs(self, start_node, destination_node_value):
    visited = set()
    queue = []
    node_steps = defaultdict(lambda: 1e7)
    node_steps[start_node] = 0
    queue.append(start_node)

    while queue:
      node = queue.pop(0)
      if node.value == destination_node_value: return node_steps[node]
      visited.add(node)

      for adjPoint in self.availablePoints(node):
        if adjPoint in visited: continue
        new_cost = node_steps[node] + 1
        if new_cost < node_steps[adjPoint]:
          node_steps[adjPoint] = new_cost
          queue.append(adjPoint)


def main():
  grid = Game()
  print(grid.bfs(grid.end, 1))
  print(grid.dijkstra(grid.end, 1))

if __name__ == "__main__":
  main()