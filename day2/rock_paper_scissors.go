package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// rock beats scissors
// scissors beats paper
// paper beats rock.
// A - X -> ROCK
// B - Y -> PAPER
// C - Z -> SCISSORS

func calculateGame(elf, me string) int {
	//rock case
	if me == "X" && elf == "A" {
		return 1 + 3
	}
	if me == "X" && elf == "B" {
		return 1
	}
	if me == "X" && elf == "C" {
		return 1 + 6
	}
	//paper case
	if me == "Y" && elf == "A" {
		return 2 + 6
	}
	if me == "Y" && elf == "B" {
		return 2 + 3
	}
	if me == "Y" && elf == "C" {
		return 2
	}

	//scissors
	if me == "Z" && elf == "A" {
		return 3
	}
	if me == "Z" && elf == "B" {
		return 3 + 6
	}
	if me == "Z" && elf == "C" {
		return 3 + 3
	}
	return 0
}

var elf string
var me string
var totalScore int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")
		elf = line[0]
		me = line[1]
		score := calculateGame(elf, me)
		totalScore += score
	}
	fmt.Println(totalScore)
}
