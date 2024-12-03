package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/PierreOudin/adventofcode/tools"
)

func main() {
	value := tools.ParseFile("../input.txt")

	r, err := regexp.Compile(`mul\((\d+\d*),(\d+\d*)\)`)

	if err != nil {
		log.Fatalln(err)
	}

	submatch := r.FindAllStringSubmatch(value, -1)

	var total int = 0

	for _, group := range submatch {
		leftNumber, _ := strconv.Atoi(group[1])
		rightNumber, _ := strconv.Atoi(group[2])
		total += leftNumber * rightNumber
	}

	fmt.Printf("Total : %v", total)
}
