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

func processGroup(group []string, score *int, sem chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	uniques := make([]string, len(group))
	for _, str := range group {
		strUniq := getUniques(str)
		uniques = append(uniques, strUniq)
	}
	groupMerged := mergeStrings(uniques)
	threeOccRune := threeOccurrencesRune(groupMerged)
	runeScore := calcCharScore(threeOccRune)
	sem <- 1
	*score += runeScore
	<-sem
}

func threeOccurrencesRune(str string) rune {
	set := stringToRuneSet(str)
	for run, app := range set {
		if app == 3 {
			return run
		}
	}
	return 0
}

func mergeStrings(arr []string) string {
	res := ""
	for _, str := range arr {
		res += str
	}
	return res
}

func getUniques(str string) string {
	set := stringToRuneSet(str)
	res := ""
	for key := range set {
		res += string(key)
	}
	return res
}

func stringToRuneSet(str string) map[rune]int {
	set := make(map[rune]int)
	for _, char := range str {
		if _, exists := set[char]; exists {
			set[char]++
		} else {
			set[char] = 1
		}
	}
	return set
}

func calcCharScore(c rune) int {
	if unicode.IsLower(c) {
		return int(c - 96)
	} else {
		return int(c - 64 + 26)
	}
}
