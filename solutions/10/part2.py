#!/usr/bin/env python3

def main():
  cycles = 0
  with open("input.txt", "r") as f:
    register = None
    x = 1
    lineIdx = 0
    display = []
    displayHeight = 6
    for i in range(displayHeight):
      display.append([0] * 40)
    while True:
      cycles += 1
      displayIdx = (cycles - 1) % 40
      if displayIdx + 1 == x or displayIdx - 1 == x or displayIdx == x:
        display[lineIdx][displayIdx] = 1
      else:
        display[lineIdx][displayIdx] = 0
      if displayIdx == 39:
        lineIdx = (lineIdx + 1) % 6

      if isinstance(register, int):
        x += register
        register = None
        continue
      line = f.readline()
      if line == "":
        print("cycles", cycles)
        break
      toks = line.strip().split(" ")
      if toks[0] == "addx":
        register = int(toks[1])
    printDisplay(display)
    
  
def printDisplay(display):
  print("--")
  for line in display:
    str = ""
    for x in line:
      if x:
        str += "#"
      else:
        str += "."
    print(str)


if __name__ == "__main__":
  main()