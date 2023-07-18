package main

import (
	"bufio"
	"fmt"
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

	arr := []int{}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		if i >= len(arr) {
			arr = append(arr, 0)
		}

		if len(line) == 0 {
			i++
		} else {
			asInt, err := strconv.Atoi(line)

			if err != nil {
				panic("Errrr")
			}

			arr[i] += asInt
		}
	}

	order := []int{0, 1, 2}

	for j := 0; j < 3; j++ {
		for i := 0; i < 3-1; i++ {
			if order[i+1] > order[i] {
				order[i], order[i+1] = order[i+1], order[i]
			}
		}
	}

	fmt.Printf("%d", order)
	println()

	k := 04
	for k < len(arr) {
		if arr[k] > arr[order[2]] {
			order[2] = k
			sort(order, arr)
		} else if arr[k] > arr[order[1]] {
			order[1] = k
			sort(order, arr)
		} else if arr[k] > arr[order[0]] {
			order[0] = k
			sort(order, arr)
		}
		k++
	}

	totalSum := arr[order[0]] + arr[order[1]] + arr[order[2]]

	fmt.Printf("%d", arr)
	println()

	println(order[0], order[1], order[2])
	println(arr[order[0]], arr[order[1]], arr[order[2]])
	println(totalSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func sort(order []int, against []int) {
	for j := 0; j < len(order); j++ {
		for i := 0; i < len(order)-1; i++ {
			if against[order[i+1]] > against[order[i]] {
				order[i], order[i+1] = order[i+1], order[i]
			}
		}
	}
}
