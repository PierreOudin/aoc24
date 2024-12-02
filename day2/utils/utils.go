package utils

import "math"

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
