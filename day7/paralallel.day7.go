package main

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic("err opening file")
	}

	println(runtime.NumGoroutine())

	defer file.Close()

	scanner := bufio.NewScanner(file)

	quant := 0

	wg := sync.WaitGroup{}
	sem := make(chan int, 8)

	start := time.Now()
	maxParallel := 1
	queue := make(chan int, maxParallel)

	for scanner.Scan() {
		wg.Add(1)
		line := scanner.Text()

		queue <- 1
		go processLine(line, &quant, &wg, sem, queue)
	}

	wg.Wait()
	duration := time.Since(start).Microseconds()

	println(quant, duration)
}

func processLine(line string, quant *int, wg *sync.WaitGroup, sem chan int, queue <-chan int) {
	defer wg.Done()

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
		sem <- 1
		*quant++
		<-sem
	}

	<-queue
}

func isInRange(initA int, endA int, initB int, endB int) bool {
	if initA >= initB && endA <= endB {
		return true
	}
	return false
}
