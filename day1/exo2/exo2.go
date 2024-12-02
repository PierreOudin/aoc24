package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PierreOudin/adventofcode/tools"
)

func main() {
	value := tools.ParseFile("../input.txt")
	t := strings.Split(value, "\n")

	var listLeft []int
	var listRight []int

	var diff int = 0
	for _, row := range t {
		if row != "" {
			rowNumbers := strings.Fields(row)

			numLeft, _ := strconv.Atoi(rowNumbers[0])
			numRight, _ := strconv.Atoi(rowNumbers[1])

			listLeft = append(listLeft, numLeft)
			listRight = append(listRight, numRight)
		}
	}

	for _, numLeft := range listLeft {
		var countLeft int = 0
		for _, num2 := range listRight {
			if numLeft == num2 {
				countLeft++
			}
		}
		diff += numLeft * countLeft
	}

	fmt.Printf("Result : %v", diff)
}
