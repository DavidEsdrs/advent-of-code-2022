package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		log.Fatal("Can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res int

	for scanner.Scan() {
		line := scanner.Text()

		value := scanLine(line)

		if value != -1 {
			res = value
			break
		}
	}

	println(res)
}

func scanLine(line string) int {
	l, r := 0, 14
	for r < len(line) {
		piece := line[l:r]
		if areUniqueCharacters(piece) {
			return r
		}
		l++
		r++
	}
	return -1
}

func areUniqueCharacters(str string) bool {
	checker := 0
	asciiA := int('a')
	for _, char := range str {
		value := int(char) - asciiA
		if checker&(1<<value) > 0 {
			return false
		}
		checker |= (1 << value)
	}
	return true
}
