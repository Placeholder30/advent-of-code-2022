import * as readline from "node:readline";
import * as fs from "fs";

const fileStream = fs.createReadStream("input.txt");

const rl = readline.createInterface({
  input: fileStream,
});
let calories = [];
let highest = 0;
let totalCalories = [];
rl.on("line", (line) => {
  if (line.match(/^$/)) {
    let sum = 0;
    for (let i = 0; i < calories.length; i++) {
      sum += calories[i];
    }
    totalCalories.push(sum);
    if (sum > highest) {
      highest = sum;
    }
    calories = [];
  }
  if (line != "") calories.push(+line);
});

rl.on("close", () => {
  let secondHighest;
  let thirdHighest;

  for (let i = 0; i < totalCalories.length; i++) {
    if (
      !secondHighest ||
      (totalCalories[i] > secondHighest && totalCalories[i] != highest)
    ) {
      secondHighest = totalCalories[i];
    }
  }
  for (let i = 0; i < totalCalories.length; i++) {
    if (
      !thirdHighest ||
      (totalCalories[i] > thirdHighest &&
        totalCalories[i] != highest &&
        totalCalories[i] != secondHighest)
    ) {
      thirdHighest = totalCalories[i];
    }
  }
  console.log({ highest, secondHighest, thirdHighest });
  console.log({ total: highest + secondHighest + thirdHighest });
});
