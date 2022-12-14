#!/usr/bin/env python3

def main():
    trees = []
    count = 0 
    with open("input.txt", "r") as f:
        for line in f:
            trees.append(list(line.strip()))
    for x in range(1, len(trees) - 1):
        for y in range(1, len(trees) - 1):
            if isVisible(x, y, trees):
                count += 1
    
    print(count)
    count += (len(trees) * 4) - 4
    print(count)

            
def isVisible(x, y, grid):
    treeValue = grid[x][y]
    # LEFT
    found = True
    for xtest in range(x):
        if grid[xtest][y] >= treeValue:
            found = False
            break
    if found:  
        print(x, y, treeValue)
        return True
    
    found = True
    # RIGHT
    for xtest in range(x + 1, len(grid)):
        if grid[xtest][y] >= treeValue:
            found = False
            break
    if found:
        print(x, y, treeValue)
        return True
    found = True
    
    # ABOVE
    for ytest in range(y):
        if grid[x][ytest] >= treeValue:
            found = False
            break
    
    if found:
        print(x, y, treeValue)
        return True
    found = True
    # BELOW
    for ytest in range(y + 1, len(grid)):
        if grid[x][ytest] >= treeValue:
            found = False
            break
    if found:
        print(x, y, treeValue)
    return found

if __name__ == "__main__":
    main()