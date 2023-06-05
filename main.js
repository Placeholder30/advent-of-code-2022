import * as readline from "node:readline";
import * as fs from "fs";

const fileStream = fs.createReadStream("input.txt");

const rl = readline.createInterface({
  input: fileStream,
});
const elfObj = {};
let calories = [];
let count = 1;
let highest = 0;
rl.on("line", (line) => {
  if (line.match(/^$/)) {
    elfObj[`elf-${count}`] = calories;
    const sum = elfObj[`elf-${count}`].reduce((acc, curr, i) => acc + curr);
    elfObj[`elf-${count}`] = sum;
    if (sum > highest) {
      highest = sum;
    }
    count++;
    calories = [];
  }
  if (line != "") calories.push(+line);
});
rl.on("close", () => {
  console.log(highest);
});
