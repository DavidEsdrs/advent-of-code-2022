package main

import (
	"bufio"
	"os"
	"time"
	"unicode"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	blockSize := 3
	group := make([]string, blockSize)
	count := 0

	score := 0

	start := time.Now()

	for scanner.Scan() {
		line := scanner.Text()
		group[count] = line
		count++
		if count >= blockSize {
			count = 0
			score += processGroup(group)
			group = make([]string, blockSize)
		}
	}

	timeEnd := time.Since(start)

	println(score, timeEnd)
}

type Tuple struct {
	item  byte
	quant int
	group int
}

func processGroup(group []string) int {
	relation := getRuneRelation(group)
	transformedRelation := transform(relation)
	mostOccurred := findMostOccurrency(transformedRelation)
	scoreForMostOccurred := calcCharScore(rune(mostOccurred.item))
	return scoreForMostOccurred
}

func getRuneRelation(group []string) map[byte]*Tuple {
	strMap := make(map[byte]*Tuple)

	for g := 0; g < len(group); g++ {
		for c := 0; c < len(group[g]); c++ {
			r := group[g][c]
			item, exists := strMap[r]

			if exists {
				if item.group != g {
					copy := strMap[r]
					copy.quant++
					copy.group = g
					strMap[r] = copy
				}
			} else {
				strMap[r] = &Tuple{item: r, quant: 1, group: g}
			}
		}
	}

	return strMap
}

func transform(input map[byte]*Tuple) []*Tuple {
	arr := []*Tuple{}
	for _, v := range input {
		arr = append(arr, v)
	}
	return arr
}

type TupleValue struct {
	item  byte
	quant int
}

func findMostOccurrency(occurrencies []*Tuple) *Tuple {
	maxIdx := 0

	for i, n := range occurrencies {
		if n.quant > occurrencies[maxIdx].quant {
			maxIdx = i
		}
	}

	return occurrencies[maxIdx]
}

func calcCharScore(c rune) int {
	if unicode.IsLower(c) {
		return int(c - 96)
	} else {
		return int(c - 64 + 26)
	}
}
