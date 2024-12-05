package main

import (
	"fmt"
	"strings"

	"github.com/PierreOudin/adventofcode/tools"
)

func main() {
	value := tools.ParseFile("../input.txt")

	horizontalRows := strings.Split(value, "\n")

	var total int = 0

	//searchedString := "MAS"

	for indexY := 0; indexY < len(horizontalRows); indexY++ {
		for indexX := 0; indexX < len(horizontalRows[indexY]); indexX++ {
			letter := string(horizontalRows[indexY][indexX])
			if letter == "A" {
				if checkXMAS(indexX, indexY, len(horizontalRows), len(horizontalRows[0]), horizontalRows) {
					total++
				}
			}
		}
	}

	fmt.Printf("Total : %v", total)
}

func checkXMAS(indexX int, indexY int, colMax int, rowMax int, rows []string) bool {
	fmt.Println("toto")
	if isValidIndex([]int{indexX, indexY}, colMax, rowMax) {
		topLeft := string(rows[indexY-1][indexX-1])
		topRight := string(rows[indexY-1][indexX+1])
		bottomLeft := string(rows[indexY+1][indexX-1])
		bottomRight := string(rows[indexY+1][indexX+1])
		if topLeft == "M" && bottomRight == "S" && topRight == "M" && bottomLeft == "S" {
			return true
		}
		if topLeft == "S" && bottomRight == "M" && topRight == "M" && bottomLeft == "S" {
			return true
		}
		if topLeft == "M" && bottomRight == "S" && topRight == "S" && bottomLeft == "M" {
			return true
		}
		if topLeft == "S" && bottomRight == "M" && topRight == "S" && bottomLeft == "M" {
			return true
		}
	}

	return false
}

func isValidIndex(index []int, colMax, rowMax int) bool {
	return index[0] > 0 && index[1] > 0 && index[0] < rowMax-1 && index[1] < colMax-1
}
