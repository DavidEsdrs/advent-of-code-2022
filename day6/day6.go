package main

import (
	"bufio"
	"os"
	"runtime"
	"sync"
	"time"
	"unicode"
)

func main() {
	runtime.GOMAXPROCS(8)
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

	sem := make(chan int, 1)

	defer close(sem)

	wg := sync.WaitGroup{}

	start := time.Now()

	for scanner.Scan() {
		line := scanner.Text()
		group[count] = line
		count++
		if count >= blockSize {
			wg.Add(1)
			go processGroup(group, &score, sem, &wg)
			count = 0
			group = make([]string, blockSize)
		}
	}

	wg.Wait()

	timeEnd := time.Since(start)
	println(score, timeEnd)
}

type Tuple struct {
	item  byte
	quant int
	group int
}

func processGroup(group []string, score *int, sem chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	relation := getRuneRelation(group)
	transformedRelation := transform(relation)
	mostOccurred := findMostOccurrency(transformedRelation)
	scoreForMostOccurred := calcCharScore(rune(mostOccurred.item))
	sem <- 1
	*score += scoreForMostOccurred
	<-sem
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
