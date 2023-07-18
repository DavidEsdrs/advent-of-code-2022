package main

import (
	"bufio"
	"os"
)

// A rock
// B paper
// C scissors

var rel = map[string]int{
	"A X": 3 + 0,
	"A Y": 1 + 3,
	"A Z": 2 + 6,

	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,

	"C X": 2 + 0,
	"C Y": 3 + 3,
	"C Z": 1 + 6,
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
