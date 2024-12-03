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

	r, err := regexp.Compile(`mul\((\d+\d*),(\d+\d*)\)|do\(\)|don't\(\)`)

	if err != nil {
		log.Fatalln(err)
	}

	submatch := r.FindAllStringSubmatch(value, -1)

	var total int = 0

	var do bool = true

	for _, group := range submatch {
		if group[0] == "don't()" {
			do = false
		} else if group[0] == "do()" {
			do = true
		} else {
			if do {
				leftNumber, _ := strconv.Atoi(group[1])
				rightNumber, _ := strconv.Atoi(group[2])
				total += leftNumber * rightNumber
			}
		}
	}

	fmt.Printf("Total : %v", total)
}
