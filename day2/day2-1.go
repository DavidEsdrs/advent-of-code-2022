package main

import (
	"bufio"
	"os"
)

var rel = map[string]int{
	"A X": 1 + 3,
	"A Y": 2 + 6,
	"A Z": 3 + 0,

	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,

	"C X": 1 + 6,
	"C Y": 2 + 0,
	"C Z": 3 + 3,
}

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

		score += rel[line]
	}

	println(score)
}
