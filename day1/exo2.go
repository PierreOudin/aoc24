package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func exo2() {
	file, err := os.Open("input.txt")
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
	t := strings.Split(value, "\n")

	var list1 []int
	var list2 []int

	var diff int = 0
	for _, row := range t {
		if row != "" {
			rowNumbers := strings.Fields(row)

			num1, err := strconv.Atoi(rowNumbers[0])
			if err != nil {
				log.Fatalf("Erreur lors de la convertion : %v", err)
			}
			num2, err := strconv.Atoi(rowNumbers[1])
			if err != nil {
				log.Fatalf("Erreur lors de la convertion : %v", err)
			}

			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}

		//fmt.Println("tesst:", row)
	}

	for _, num1 := range list1 {
		var countNum1 int = 0
		for _, num2 := range list2 {
			if num1 == num2 {
				countNum1++
			}
		}
		diff += num1 * countNum1
	}

	fmt.Printf("Result : %v", diff)
}
