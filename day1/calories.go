package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(data)
	lineSlice := []int{}
	sum := []string{}
	var highest int
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			var total int
			for _, item := range sum {
				item, _ := strconv.Atoi(item)
				total += item
			}
			if total > highest {
				highest = total
			}
			lineSlice = append(lineSlice, total)
			sum = nil
			continue
		}
		sum = append(sum, (line))
	}
	var secondHighest, thirdHighest int
	for _, item := range lineSlice {
		if item != highest && item > secondHighest {
			secondHighest = item
		}
	}
	for _, item := range lineSlice {
		if item != highest && item != secondHighest && item > thirdHighest {
			thirdHighest = item
		}
	}

	fmt.Println(highest + secondHighest + thirdHighest)
}
