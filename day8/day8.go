package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	quant := 0
	start := time.Now()

	for scanner.Scan() {
		line := scanner.Text()

		elves := strings.Split(line, ",")

		firstAssg := elves[0]
		startEnd1 := strings.Split(firstAssg, "-")
		init1, _ := strconv.Atoi(startEnd1[0])
		end1, _ := strconv.Atoi(startEnd1[1])

		secondAssg := elves[1]
		startEnd2 := strings.Split(secondAssg, "-")
		init2, _ := strconv.Atoi(startEnd2[0])
		end2, _ := strconv.Atoi(startEnd2[1])

		if isInRange(init1, end1, init2, end2) || isInRange(init2, end2, init1, end1) {
			quant++
		}
	}

	elapsed := time.Since(start)
	println(quant, elapsed)
}

func isInRange(initA int, endA int, initB int, endB int) bool {
	if initA >= initB && initA <= endB {
		return true
	} else if endA >= initB && endA <= endB {
		return true
	}
	return false
}
