def main():
    top3 = [0, 0, 0]
    curr_sum = 0
    with open("input.txt", "r") as f:
        for line in f:
            val = line.strip()
            if val == "":
                if curr_sum > top3[2]:
                    if curr_sum > top3[1]:
                        if curr_sum > top3[0]:
                            top3 = [curr_sum, top3[0], top3[1]]
                        else:
                            top3 = [top3[0], curr_sum, top3[1]]
                    else:
                        top3[2] = curr_sum
                curr_sum = 0
            else:
                curr_sum += int(val)
    print(sum(top3))


if __name__ == "__main__":
    main()