import * as fs from 'fs';

const input = fs.readFileSync("input.txt", "utf-8")
const lines = input.split("\n")

const totals: number[] = []
let currSum = 0
for (let line of lines) {
    if (line == "") {
        totals.push(currSum)
        currSum = 0
    } else {
        currSum += Number(line)
    }
}
totals.sort().reverse()
console.log("Part 1:", totals[0])
console.log("Part 2:", totals[0] + totals[1] + totals[2])


