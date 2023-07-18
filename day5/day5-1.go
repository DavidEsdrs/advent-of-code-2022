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

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr := make([][]string, 9)

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

		top := arr[from][:quant]

		println(quant, from, to)

		fmt.Printf("before: %v, from = %v", arr[from], from+1)
		println()

		arr[from] = arr[from][quant:]

		fmt.Printf("after: %v", arr[from])
		println()

		fmt.Printf("dest before: %+v, to = %v", arr[to], to+1)
		println()

		for _, n := range top {
			arr[to] = append([]string{n}, arr[to]...)
		}

		fmt.Printf("dest after: %v", arr[to])
		println()
		println()
	}

	str := ""

	for _, n := range arr {
		str += n[0]
		fmt.Printf("%v", n)
		println()
	}
	println()

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
