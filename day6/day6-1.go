package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res int

	for scanner.Scan() {
		line := scanner.Text()

		index, found := scanLine(line)

		if found {
			res = index
			break
		}
	}

	println(res)
}

func scanLine(line string) (index int, found bool) {
	l, r := 0, 4
	for r < len(line) {
		part := line[l:r]

		if checkUnique(part) {
			return r, true
		}

		l++
		r++
	}

	return -1, false
}

func checkUnique(str string) bool {
	checker := 0
	for _, char := range str {
		value := int(char) - int('a')
		if (checker & (1 << value)) > 0 {
			return false
		}
		checker |= (1 << value)
	}
	return true
}
