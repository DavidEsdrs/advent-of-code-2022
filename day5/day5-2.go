package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cratesQuant int

func init() {
	args := os.Args[1:]
	argAsNum, err := strconv.Atoi(args[0])
	if err != nil {
		panic("Invalid quant")
	}
	cratesQuant = argAsNum
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr := make([][]string, cratesQuant)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			for f := 0; f < len(line)+1; f += 4 {
				crate := line[f : f+3]
				// if it is in fact a crate, not something else
				if len(strings.TrimSpace(crate)) >= 3 {
					arr[f/4] = append(arr[f/4], crate[1:2])
				}
			}
		} else {
			break
		}
	}

	_print(arr)

	regex := regexp.MustCompile("[^0-9]+|\\s+")

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		output := regex.ReplaceAllString(line, " ")
		output = strings.TrimSpace(output)
		outputArr := strings.Split(output, " ")

		quant := getInt(outputArr[0])
		from := getInt(outputArr[1]) - 1
		to := getInt(outputArr[2]) - 1

		top := make([]string, quant)
		copy(top, arr[from][:quant])

		arr[from] = arr[from][quant:]

		arr[to] = append(top, arr[to]...)
	}

	str := ""

	for _, n := range arr {
		if len(n) > 0 {
			str += n[0]
		}
	}

	_print(arr)

	println(str)

}

func getInt(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}

func reverse(array []string) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func _print(array [][]string) {
	println("======================================")
	for i, n := range array {
		fmt.Printf("%v %v", n, i+1)
		println()
	}
	println("======================================")
}
