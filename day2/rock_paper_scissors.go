package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateGame(elf, you string) int {
	//loss cases
	if you == "X" && elf == "A" {
		return 3
	}
	if you == "X" && elf == "B" {
		return 1
	}
	if you == "X" && elf == "C" {
		return 2
	}
	//draw cases
	if you == "Y" && elf == "A" {
		return 1 + 3
	}
	if you == "Y" && elf == "B" {
		return 2 + 3
	}
	if you == "Y" && elf == "C" {
		return 3 + 3
	}

	//win cases
	if you == "Z" && elf == "A" {
		return 2 + 6
	}
	if you == "Z" && elf == "B" {
		return 3 + 6
	}
	if you == "Z" && elf == "C" {
		return 1 + 6
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	var elf string
	var you string
	var totalScore int
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")
		elf, you = line[0], line[1]

		score := calculateGame(elf, you)
		totalScore += score
	}
	fmt.Println(totalScore)
}
