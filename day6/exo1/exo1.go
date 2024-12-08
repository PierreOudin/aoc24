package main

import (
	"fmt"
	"strings"

	"github.com/PierreOudin/adventofcode/tools"
)

type Coordinate struct {
	X int
	Y int
}

func main() {
	value := tools.ParseFile("../input.txt")

	mapLines := strings.Split(value, "\n")

	directions := make(map[rune]Coordinate, 0)
	directions['^'] = Coordinate{X: 0, Y: -1}
	directions['<'] = Coordinate{X: -1, Y: 0}
	directions['>'] = Coordinate{X: 1, Y: 0}
	directions['v'] = Coordinate{X: 0, Y: 1}

	guardStep := 0

	isInside := true

	for isInside {
		for y := 0; y < len(mapLines); y++ {
			line := mapLines[y]
			for x := 0; x < len(line); x++ {
				if checkGuardCarac(rune(line[x])) {
					carac := rune(line[x])
					nextIndex := Coordinate{X: x + directions[carac].X, Y: y + directions[carac].Y}
					if nextIndex.X < 0 || nextIndex.X >= len(line) || nextIndex.Y < 0 || nextIndex.Y >= len(mapLines) {
						isInside = false
						guardStep++
					} else {
						nextCarac := rune(mapLines[nextIndex.Y][nextIndex.X])
						if nextCarac == '.' || nextCarac == 'X' {
							if nextIndex.Y != y {
								nextLine := mapLines[nextIndex.Y]
								nextLine = replaceAtIndex(nextLine, carac, nextIndex.X)
								mapLines[nextIndex.Y] = nextLine
							} else {
								line = replaceAtIndex(line, carac, nextIndex.X)
							}
							line = replaceAtIndex(line, 'X', x)
							mapLines[y] = line
							if nextCarac != 'X' {
								guardStep++
							}
						} else if nextCarac == '#' {
							carac = ChangeDir(carac)
							line = replaceAtIndex(line, carac, x)
							mapLines[y] = line
						}
					}

					fmt.Println(guardStep)
					PrintTab(mapLines)
				}
			}

		}
	}

	fmt.Printf("Steps : %v", guardStep)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func ChangeDir(carac rune) rune {
	switch carac {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}
	return 'X'
}

func PrintTab(tab []string) {
	fmt.Println("-----")
	for _, t := range tab {
		fmt.Println(t)
	}
	fmt.Println("-----")
}

func checkGuardCarac(carac rune) bool {
	return carac == '^' || carac == '<' || carac == '>' || carac == 'v'
}
