def main():
    max = 0 
    curr_sum = 0
    with open("input.txt", "r") as f:
        for line in f:
            val = line.strip()
            if val == "":
                if curr_sum > max:
                    max = curr_sum
                curr_sum = 0
            else:
                curr_sum += int(val)
    print(max)


if __name__ == "__main__":
    main()