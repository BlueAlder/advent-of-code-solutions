#!/usr/bin/env python3

def main():
    totals = []
    curr_sum = 0
    with open("input.txt", "r") as f:
        for line in f:
            val = line.strip()
            if val == "":
                totals.append(curr_sum)
                curr_sum = 0
            else:
                curr_sum += int(val)
    print("Part 1:", max(totals))
    totals.sort(reverse=True)
    print("Part 2:", totals[0] + totals[1] + totals[2])


if __name__ == "__main__":
    main()