package main

import (
	"bufio"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		mid := len(line) / 2

		score += findAInB(line[:mid], line[mid:])
	}
	println(score)
}

func findAInB(strA string, strB string) int {
	for _, charA := range strA {
		for _, charB := range strB {
			if charA == charB {
				if unicode.IsLower(rune(charB)) {
					return int(charB - 96)
				} else {
					return int(charB - 64 + 26)
				}
			}
		}
	}

	return 0
}
