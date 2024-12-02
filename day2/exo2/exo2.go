package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture du fichier : %v", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	f, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
	}
	value := string(f)

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
		//fmt.Println(line)
		if IsLineOnlyAscOrDesc(line) && IsLineOnlyInRange(line) {
			//fmt.Println("Toto")
			numSafeRow++
		} else {
			var testingLines [][]int
			for index := 0; index < len(line); index++ {
				removeIndexLine := RemoveIndex(line, index)
				testingLines = append(testingLines, removeIndexLine)
			}
			fmt.Println(testingLines)
			for _, testingLine := range testingLines {
				if IsLineOnlyAscOrDesc(testingLine) && IsLineOnlyInRange(testingLine) {
					numSafeRow++
					break
				}
			}
		}
	}

	fmt.Printf("Row safe : %v", numSafeRow)
}

func IsLineOnlyAscOrDesc(line []int) bool {
	var asc bool = true
	var desc bool = true
	for index := 0; index < len(line)-1; index++ {
		asc = asc && (line[index] < line[index+1])
		desc = desc && (line[index] > line[index+1])
	}
	return asc || desc
}

func IsLineOnlyInRange(line []int) bool {
	var inRange bool = true
	for index := 0; index < len(line)-1; index++ {
		var diff int = int(math.Abs(float64(line[index] - line[index+1])))
		if !(diff >= 1 && diff <= 3) {
			inRange = false
			break
		}
	}
	return inRange
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
