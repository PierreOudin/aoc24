package tools

import (
	"io"
	"log"
	"os"
)

func tools() {

}

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

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
