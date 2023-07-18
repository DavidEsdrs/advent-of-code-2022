package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr := [][]int{}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		if i >= len(arr) {
			arr = append(arr, []int{})
		}

		if len(line) == 0 {
			i++

		} else {
			asInt, err := strconv.Atoi(line)

			if err != nil {
				panic("Errrr")
			}
			arr[i] = append(arr[i], asInt)
		}
	}

	maxes := make([]int, len(arr))

	for i, n := range arr {
		maxes[i] = _map(n)
	}

	maxIdx := 0

	for i, n := range maxes {
		if n > maxes[maxIdx] {
			maxIdx = i
		}
	}

	println(maxes[maxIdx])

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func _map(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}
