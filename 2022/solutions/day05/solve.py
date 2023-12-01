#!/usr/bin/env python3

NUM_STACKS = 9

def main(part):
    stacks = parseInput()
    with open("input.txt", "r") as f:
        for line in f:
            if line[0] == "m":
                tokens = line.split(" ")
                amount = int(tokens[1])
                from_stack = int(tokens[3]) - 1
                to_stack = int(tokens[5]) - 1
                
                if part == 1:
                    for i in range(amount):
                        move = stacks[from_stack].pop()
                        stacks[to_stack].append(move)
                elif part == 2:
                    to_move = []
                    for i in range(amount):
                        to_move.append(stacks[from_stack].pop())
                    to_move_ordered = to_move[::-1]
                    stacks[to_stack] = stacks[to_stack] + to_move_ordered 

    join = ""
    for stack in stacks:
        join += stack.pop()
    return join




def parseInput():
    stacks = [[],[],[],[],[],[],[],[],[]]
    with open("input.txt", "r") as f:
        lines = []
        for line in f:
            if len(line) == 36:
                n = 4
                a = [line[i:i + n] for i in range(0, len(line), n)]
                lines.append(a)
        lines.pop()
        lines = list(zip(*lines[::-1]))
        for idx, stack in enumerate(lines):
            for item in stack:
                if item[1] != " ":
                    # print(stacks[idx])
                    stacks[idx].append(item[1])
    
    return stacks


if __name__ == "__main__":
    print(main(1))
    print(main(2))