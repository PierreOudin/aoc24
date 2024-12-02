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

	var numIsSafe int = 0

	for _, line := range formattedLines {
		//fmt.Println(line)
		if utils.IsLineOnlyAscOrDesc(line) && utils.IsLineOnlyInRange(line) {
			//fmt.Println("Toto")
			numIsSafe++
		}
	}

	fmt.Println(numIsSafe)
	fmt.Println()
}
