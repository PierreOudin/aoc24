package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func exo1() {
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

	sort.Ints(list1)
	sort.Ints(list2)

	fmt.Println(list2)

	for i := 0; i < len(list1); i++ {
		diff += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Printf("Result : %v", diff)
}
