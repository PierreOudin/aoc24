package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PierreOudin/adventofcode/tools"
)

type Line struct {
	result  int
	numbers []int
}

func main() {
	value := tools.ParseFile("../input.txt")

	mapLines := strings.Split(value, "\n")

	splitedLines := make([]Line, 0)

	for _, mapLine := range mapLines {
		split := strings.Split(mapLine, ":")
		result, _ := strconv.Atoi(split[0])
		listNum := make([]int, 0)
		for _, num := range strings.Split(split[1], " ") {
			if num != "" {
				number, _ := strconv.Atoi(num)
				listNum = append(listNum, number)
			}
		}
		splitedLines = append(splitedLines, Line{result: result, numbers: listNum})
	}

	total := 0

	for _, line := range splitedLines {
		if EquationWork(line, line.numbers[0], 1) {
			total += line.result
		}
	}

	fmt.Printf("Total : %v", total)
}

func EquationWork(line Line, sum, index int) bool {
	if sum > line.result {
		return false
	}
	if len(line.numbers) <= index {
		return line.result == sum
	}
	return EquationWork(line, sum+line.numbers[index], index+1) ||
		EquationWork(line, sum*line.numbers[index], index+1) ||
		EquationWork(line, Concat(sum, line.numbers[index]), index+1)
}

func Concat(i, j int) int {
	concat, _ := strconv.Atoi(strconv.Itoa(i) + strconv.Itoa(j))
	fmt.Println(concat)
	return concat
}
