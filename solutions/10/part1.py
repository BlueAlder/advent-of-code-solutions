#!/usr/bin/env python3

def main():
  cycles = 0
  total = 0
  with open("input.txt", "r") as f:
    register = None
    x = 1
    while True:
      cycles += 1
      if cycles in [20, 60, 100, 140, 180, 220]:
        signal = x * cycles
        print(cycles, ":", signal)
        total += signal
      if isinstance(register, int):
        x += register
        register = None
        continue
      line = f.readline()
      if line == "":
        break
      toks = line.strip().split(" ")
      if toks[0] == "addx":
        register = int(toks[1])
  print(total)


if __name__ == "__main__":
  main()