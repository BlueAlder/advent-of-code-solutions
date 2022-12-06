#!/usr/bin/env python3 

def main():
    file = open("input.txt", "r")
    line = file.readline()

    message_len = 4
    
    window = list(line[:message_len])
    print(window)
    char_count = message_len

    for i in range(message_len - 1, len(line)):
        char = line[i]
        window.pop(0)
        window.append(char)
        print(window)
        if len(window) == len(set(window)):
            break
        char_count += 1
    print(char_count)



if __name__ == "__main__":
    main()