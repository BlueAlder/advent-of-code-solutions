#!/usr/bin/env python3 

def solve(message_len):
    file = open("input.txt", "r")
    line = file.readline()

    window = list(line[:message_len])
    char_count = message_len

    for i in range(message_len - 1, len(line)):
        char = line[i]
        window.pop(0)
        window.append(char)
        if len(window) == len(set(window)):
            break
        char_count += 1
    return char_count

def main():
    print("Part 1:", solve(4))
    print("Part 2:", solve(14))


if __name__ == "__main__":
    main()