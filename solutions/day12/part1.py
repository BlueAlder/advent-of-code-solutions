#!/usr/bin/env python3 


from collections import defaultdict
import heapq as heap

class Point:
  def __init__(self, x, y, value):
    self.x = x
    self.y = y
    self.value = value
  
  def __str__(self):
    return f"({self.x},{self.y})"

class Game:
  def __init__(self):
    grid = []
    start = []
    end = []
    with open("/home/sam/Documents/advent-of-code-22/solutions/12/input.txt", "r") as f:
      y = 0
      for line in f:
        row = []
        for idx, chr in enumerate(line.strip()):
          if chr == "S":
            start = Point(idx, y, 1)
            row.append(start)
          elif chr == "E":
            end = Point(idx, y, 26)
            row.append(end)
          else:
            row.append(Point(idx, y, ord(chr) - 96))

        y += 1
        grid.append(row)
    self.grid = grid
    self.start = start
    self.end = end

  @property
  def width(self):
    return len(self.grid[0])
  
  @property
  def height(self):
    return len(self.grid)


  def availablePoints(self, current_point):
    point_val = current_point.value
    points = []

    if current_point.y - 1 >= 0 and self.grid[current_point.y - 1][current_point.x].value - 1 <= point_val :
      points.append(self.grid[current_point.y - 1][current_point.x])
    if current_point.y + 1 < self.height  and self.grid[current_point.y + 1][current_point.x].value - 1 <= point_val:
      points.append(self.grid[current_point.y + 1][current_point.x])
    if current_point.x - 1 >= 0  and self.grid[current_point.y][current_point.x - 1].value - 1 <= point_val:
      points.append(self.grid[current_point.y][current_point.x - 1])
    if current_point.x + 1 < self.width  and self.grid[current_point.y][current_point.x + 1].value - 1 <= point_val: 
      points.append(self.grid[current_point.y][current_point.x + 1])
    
    return points


    
  def dijkstra(self, start_node):
    visited = set()
    parentsMap = {}
    pq = []
    node_costs = defaultdict(lambda: 1e7)
    node_costs[start_node] = 0
    entry_count = 0
    heap.heappush(pq, (0, entry_count, start_node))

    while pq:
      _, _, node = heap.heappop(pq)
      visited.add(node)

      for point in self.availablePoints(node):
        if point in visited: continue

        entry_count += 1
        new_cost = node_costs[node] + 1
        if node_costs[point] > new_cost:
          parentsMap[point] = node
          node_costs[point] = new_cost
          heap.heappush(pq, (new_cost, entry_count, point))
    return node_costs[self.end]



def main():
  grid = Game()
  val = grid.dijkstra(grid.start)
  print(val)



if __name__ == "__main__":
  main()