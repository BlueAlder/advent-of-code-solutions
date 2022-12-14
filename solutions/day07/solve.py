#!/usr/bin/env python3 

def main():
    current_dir = "/"
    dir_sizes = {}
    with open("input.txt", "r") as f:
        for line in f:
            if line[0] == "$":
                tokens = line.split(' ')
                if tokens[1] == "cd":
                    current_dir = change_dir(current_dir, tokens[2].strip())
            else:
                file_toks = line.split(" ")
                if file_toks[0] != "dir":
                    file_size = int(file_toks[0])
                    tmp_dir = current_dir
                    while True:
                        if tmp_dir not in dir_sizes:
                            dir_sizes[tmp_dir] = file_size
                        else:
                            dir_sizes[tmp_dir] += file_size
                        if tmp_dir == "/":
                            break
                        tmp_dir = pop_dir(tmp_dir)
    count = 0
    for idx, (dir, val) in enumerate(dir_sizes.items()):
        if val < 100000:
            count += val
    print("Part 1:", count)

    total_space = 70000000
    required_space =  30000000                  

    unused_space =  total_space - dir_sizes["/"]
    dir_size_min = required_space - unused_space
    closest = total_space

    for idx, (dir, val) in enumerate(dir_sizes.items()):
        if val > dir_size_min and val < closest:
            closest = val
    print("Part 2:", closest)


def change_dir(current_dir, arg):
    if arg[0] == "/":
        return arg
    if arg == "..":
        return pop_dir(current_dir)
    if current_dir == "/":
        return current_dir + arg
    else:
        return current_dir + "/" + arg

def pop_dir(current_dir):
    folders = current_dir.split("/")
    folders.pop()
    if len(folders) == 1:
        return "/"
    return '/'.join(folders)


if __name__ == "__main__":
    main()