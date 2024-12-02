package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PierreOudin/adventofcode/day2/utils"
	"github.com/PierreOudin/adventofcode/tools"
)

func main() {
	value := tools.ParseFile("../input.txt")

	var formattedLines [][]int

	for _, row := range strings.Split(value, "\n") {
		var formattedLine []int
		for _, field := range strings.Fields(row) {
			fieldInt, _ := strconv.Atoi(field)
			formattedLine = append(formattedLine, fieldInt)
		}
		formattedLines = append(formattedLines, formattedLine)
	}

	var numSafeRow int = 0

	for _, line := range formattedLines {
		if utils.IsLineOnlyAscOrDesc(line) && utils.IsLineOnlyInRange(line) {
			numSafeRow++
		} else {
			var testingLines [][]int
			for index := 0; index < len(line); index++ {
				removeIndexLine := tools.RemoveIndex(line, index)
				testingLines = append(testingLines, removeIndexLine)
			}
			for _, testingLine := range testingLines {
				if utils.IsLineOnlyAscOrDesc(testingLine) && utils.IsLineOnlyInRange(testingLine) {
					numSafeRow++
					break
				}
			}
		}
	}

	fmt.Printf("Row safe : %v", numSafeRow)
}
