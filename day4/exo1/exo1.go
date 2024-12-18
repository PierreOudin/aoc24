package main

import (
	"fmt"
	"strings"

	"github.com/PierreOudin/adventofcode/tools"
)

func main() {
	value := tools.ParseFile("../input.txt")

	directions := [8][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, -1},
		{1, 1},
		{-1, -1},
		{-1, 1},
	}
	horizontalRows := strings.Split(value, "\n")

	var total int = 0

	searchedString := "XMAS"

	for indexY := 0; indexY < len(horizontalRows); indexY++ {
		for indexX := 0; indexX < len(horizontalRows[indexY]); indexX++ {
			letter := string(horizontalRows[indexY][indexX])
			if letter == string(searchedString[0]) {
				for _, direction := range directions {
					checkXMAS(indexX, indexY, direction, searchedString, 0, &total, len(horizontalRows), len(horizontalRows[0]), horizontalRows)
				}
			}
		}
	}

	fmt.Printf("Total : %v", total)
}

func checkXMAS(indexX int, indexY int, direction [2]int, searchedWord string, indexWord int, total *int, colMax int, rowMax int, rows []string) {
	if indexWord+1 == len(searchedWord) {
		(*total)++
		return
	}

	newIndex := []int{indexX + direction[0], indexY + direction[1]}
	if isValidIndex(newIndex, colMax, rowMax) {
		nextLetter := string(rows[newIndex[1]][newIndex[0]])
		if nextLetter == string(searchedWord[indexWord+1]) {
			checkXMAS(newIndex[0], newIndex[1], direction, searchedWord, indexWord+1, total, colMax, rowMax, rows)
		}
	}
}

func isValidIndex(index []int, colMax, rowMax int) bool {
	return index[0] >= 0 && index[1] >= 0 && index[0] < rowMax && index[1] < colMax
}
