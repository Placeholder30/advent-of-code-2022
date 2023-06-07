package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var sum int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		seperator := int((len(line)) / 2)
		firstCompartment, secondCompartment := line[:seperator], line[seperator:]
		var match string
		for i := 0; i < len(firstCompartment); i++ {
			if match != "" {
				match = ""
				break
			}
			for j := 0; j < len(secondCompartment); j++ {
				if firstCompartment[i] == secondCompartment[j] {
					fmt.Println(firstCompartment[i])
					match = (string(firstCompartment[i]))
					alphabetMap := AlphabetMap[match]
					sum = sum + alphabetMap
					break
				}
			}

		}

	}
	fmt.Println(sum)
}
