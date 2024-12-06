package tools

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ParseFile(path string) string {
	file, err := os.Open(path)
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
	return string(f)
}

func ParseFileScanner(path string) *bufio.Scanner {
	file, err := os.Open("../input.txt")
	//file, err := os.Open("../example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return bufio.NewScanner(file)
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
