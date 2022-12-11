#!/usr/bin/env python3

class Monkey:
  def __init__(self, items, operation, divisibilityTest, true, false):
    self.items = items
    self.operation = operation
    self.divisibilityTest = divisibilityTest
    self.true = true
    self.false = false
    self.inspected = 0


  def inspectItem(self, idx):
    self.inspected += 1
    item = self.items[idx]
    newItem = self.operation(item)
    newItem = newItem % 9699690
    if newItem % self.divisibilityTest == 0:
      return [newItem, self.true]
    else:
      return [newItem, self.false]
  
  def __str__(self):
    return str(self.items)

def main():
  monkeys = loadMonkeys()
  round_count = 10000
  for i in range(round_count):
    playRound(monkeys)
    print('round', i)
  
  inspections = []
  for monkey in monkeys:
    inspections.append(monkey.inspected)
  inspections.sort()
  mb = inspections[-1] * inspections[-2]
  print(inspections)
  for idx, monkey in enumerate(monkeys):
    print(idx, ":", monkey.inspected, monkey)
  print("Monkey Business", mb)

def playRound(monkeys):
  for monkey in monkeys:
    for idx in range(len(monkey.items)):
      val = monkey.inspectItem(idx)
      monkeys[val[1]].items.append(val[0])
    monkey.items = []

def loadMonkeys():
  monkeys = []
  with open("/home/sam/Documents/advent-of-code-22/solutions/11/input.txt", "r") as f:
    while True:
      f.readline()
      itemstr = f.readline().strip().split(" ")[2::]
      items = []
      for i in itemstr:
        items.append(int(i.replace(",", "").strip()))

      operationstr = f.readline().strip().split(" ")[-2:]
      if operationstr[1] == "old":
        operationstr[1] = "x"
      evalstr = "x" + operationstr[0] + operationstr[1]
      operation = eval("lambda x: "+evalstr)
      div = int(f.readline().strip().split(" ")[-1])
      true = int(f.readline().strip().split(" ")[-1])
      false = int(f.readline().strip().split(" ")[-1])

      monkeys.append(Monkey(items, operation, div, true, false))
      if f.readline() == "":
        break
  return monkeys

if __name__ == "__main__":
  main()