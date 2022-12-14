#!/usr/bin/env python3

def main():
    trees = loadGrid()
    score_max =  0
    
    for x in range(1, len(trees) - 1):
        for y in range(1, len(trees) - 1):
            score = viewScore(x, y, trees)
            if score > score_max:
                score_max = score
    
    print(score_max)

def loadGrid():
    trees = []
    with open("input.txt", "r") as f:
        for line in f:
            trees.append(list(line.strip()))
    return trees
            
def viewScore(x, y, grid):
    treeValue = grid[x][y]

    # LEFT
    left = 0
    for xtest in range(x - 1, -1, -1):
        left += 1
        if grid[xtest][y] >= treeValue:
            break
    
    right = 0
    # RIGHT
    for xtest in range(x + 1, len(grid)):
        right += 1
        if grid[xtest][y] >= treeValue:
            break
    
    # ABOVE
    above = 0
    for ytest in range(y - 1, -1 ,-1):
        above += 1
        if grid[x][ytest] >= treeValue:
            break
    
    # BELOW
    below = 0
    for ytest in range(y + 1, len(grid)):
        below += 1
        if grid[x][ytest] >= treeValue:
            break
    return left * right * above * below

if __name__ == "__main__":
    main()