#!/usr/bin/env python3

import itertools

def main():
    with open("/usr/local/google/home/samcalamos/Documents/personal/advent-of-code-22/solutions/13/input.txt", "r") as f:
        input = f.read()
    pairs = input.split("\n\n")

    total = 0
    for idx, pair in enumerate(pairs):

        [left, right] = pair.split("\n")
        left = eval(left)
        right = eval(right)
    
        if comparePackets(left, right): total += idx + 1
    print(total)
        
def comparePackets(left, right):
    for l, r in itertools.zip_longest(left, right): 
        if l == None or r == None:
            return l == None
        if isinstance(l, int) and isinstance(r, int):
            if l < r:
                return True
            elif l > r:
                return False
            else:
                continue
        else:
            res = comparePackets(convertToList(l), convertToList(r)) 
            if res != None:
                return res
    
def convertToList(input):
    if isinstance(input, int):
        return [input]
    return input

if __name__ == "__main__":
    main()