package main

import (
	"fmt"
	"math"
	"sort"
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

	sort.Ints(listLeft)
	sort.Ints(listRight)

	for i := 0; i < len(listLeft); i++ {
		diff += int(math.Abs(float64(listLeft[i] - listRight[i])))
	}

	fmt.Printf("Result : %v", diff)
}
