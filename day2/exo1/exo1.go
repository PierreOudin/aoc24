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
	//t := strings.Split(value, "\n")

	//var numSafe int = 0

	var numIsSafe int = 0

	for _, row := range strings.Split(value, "\n") {
		fieldArray := strings.Fields(row)

		var isAsc bool = false

		for i := 0; i < len(fieldArray)-1; i++ {
			//fmt.Println(i)
			num, err := strconv.Atoi(fieldArray[i])
			if err != nil {
				log.Fatalf("Erreur lors de la convertion : %v", err)
			}
			nextNum, err := strconv.Atoi(fieldArray[i+1])
			if err != nil {
				log.Fatalf("Erreur lors de la convertion : %v", err)
			}
			diff := int(math.Abs(float64(num - nextNum)))
			fmt.Println(diff)
			if diff >= 1 && diff <= 3 {
				if i == 0 {
					isAsc = num < nextNum
				} else {
					if (num < nextNum && !isAsc) || (num > nextNum && isAsc) {
						break
					}
				}

				if i == len(fieldArray)-2 {
					fmt.Println(fieldArray)
					numIsSafe++
				}
			} else {
				break
			}

		}
	}

	fmt.Println(numIsSafe)
	fmt.Println()
}
