package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/PierreOudin/adventofcode/tools"
)

func main() {
	value := tools.ParseFile("../input.txt")

	var tabRules []string
	var tabPrint []string

	firstPartOver := false

	for _, line := range strings.Split(value, "\n") {
		if line == "" {
			fmt.Println("toto")
			firstPartOver = true
			continue
		} else {
			if !firstPartOver {
				tabRules = append(tabRules, line)
			} else {
				tabPrint = append(tabPrint, line)
			}
		}
	}

	rules := make(map[int][]int)

	for _, rule := range tabRules {
		xy := strings.Split(rule, "|")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		rules[x] = append(rules[x], y)
	}

	result := make([][]int, 0)

	for _, print := range tabPrint {
		printNumbers := strings.Split(print, ",")
		numbers := make([]int, 0)
		for _, number := range printNumbers {
			num, _ := strconv.Atoi(number)
			numbers = append(numbers, num)
		}

		if !verifyRule(numbers, rules) {
			result = append(result, numbers)
		}
	}

	fmt.Println(result)

	finalLines := make([][]int, 0)

	for _, wrongLine := range result {
		finalLines = append(finalLines, sortPrintLine(wrongLine, rules))
	}

	fmt.Println(finalLines)

	total := 0

	for _, res := range result {
		total += res[int(len(res)/2)]
	}

	fmt.Printf("Total : %v", total)
}

func verifyRule(line []int, rules map[int][]int) bool {
	for index := 0; index < len(line); index++ {
		rule := rules[line[index]]
		for _, r := range rule {
			ruleIndex := slices.Index(line, r)
			if ruleIndex != -1 && ruleIndex < index {
				return false
			}
		}
	}
	return true

}

func sortPrintLine(line []int, rules map[int][]int) []int {
	for !verifyRule(line, rules) {
		for index := 0; index < len(line); index++ {
			rule := rules[line[index]]
			for _, r := range rule {
				ruleIndex := slices.Index(line, r)
				if ruleIndex != -1 && ruleIndex < index {
					line[index], line[ruleIndex] = line[ruleIndex], line[index]
				}
			}
		}
	}
	return line
}
