#!/usr/bin/env python3

import itertools
import functools

def main():
    with open("/usr/local/google/home/samcalamos/Documents/personal/advent-of-code-22/solutions/13/input.txt", "r") as f:
        input = f.read()
    pairs = input.split("\n")
    allElements = list(filter(lambda e: e != "", pairs))
    allElements = list(map(lambda e: eval(e), allElements))
    allElements.append([[2]])
    allElements.append([[6]])

    sorted_els = sorted(allElements, key=functools.cmp_to_key(comparePackets))

    i1 = sorted_els.index([[2]]) + 1
    i2 = sorted_els.index([[6]]) + 1

    print(i1 * i2)
        
def comparePackets(left, right):
    for l, r in itertools.zip_longest(left, right): 
        if l == None or r == None:
            if l == None: return -1
            return 1
        if isinstance(l, int) and isinstance(r, int):
            if l < r:
                return -1
            elif l > r:
                return 1
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