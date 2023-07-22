package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		log.Fatal("Can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rows := 99
	cols := 99

	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	column := 0
	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, char := range line {
			a := string(char)
			value, _ := strconv.Atoi(a)
			matrix[row][column] = value
			column++
		}

		column = 0
		row++
	}

	visible := (2 * rows) + (2 * cols) - 4

	start := time.Now()

	for i, j := 1, 1; i < len(matrix)-1; {
		if !HiddenTop(matrix, i, j) ||
			!HiddenLeft(matrix[i], i, j) ||
			!HiddenBottom(matrix, i, j) ||
			!HiddenRight(matrix[i], i, j) {
			visible++
		}

		if j == len(matrix[i])-2 {
			j = 1
			i++
		} else {
			j++
		}

	}

	duration := time.Since(start).Milliseconds()

	println("It took:", duration, "milliseconds")

	fmt.Printf("%v", visible)
}

func HiddenTop(matrix [][]int, i int, j int) bool {
	col := GetCol(matrix, j, 0, i)

	height := matrix[i][j]

	for _, h := range col {
		if h >= height {
			return true
		}
	}

	return false
}

func HiddenBottom(matrix [][]int, i int, j int) bool {
	height := matrix[i][j]

	col := GetCol(matrix, j, i+1, len(matrix[i]))

	for _, h := range col {
		if h >= height {
			return true
		}
	}

	return false
}

func GetCol(matrix [][]int, col int, indexes ...int) []int {
	if len(indexes) == 0 {

		arr := make([]int, len(matrix[0]))

		for i := 0; i < len(matrix[0]); i++ {
			arr[i] = matrix[i][col]
		}

		return arr

	} else if len(indexes) == 1 {

		startIndex := indexes[0]

		length := len(matrix) - startIndex

		arr := make([]int, length)

		for i := indexes[0]; i < len(matrix); i++ {
			arr[i-startIndex] = matrix[i][col]
		}

		return arr

	} else if len(indexes) == 2 {

		startIndex := indexes[0]
		endIndex := indexes[1]

		length := endIndex - startIndex

		if endIndex > len(matrix) {
			log.Fatal("The end index is greater than the matrix row length")
		}

		arr := make([]int, length)

		for i := startIndex; i < endIndex; i++ {
			arr[i-startIndex] = matrix[i][col]
		}

		return arr

	}
	panic("Too many arguments")
}

func HiddenRight(row []int, i int, j int) bool {
	value := row[j]
	for i := j + 1; i < len(row); i++ {
		if value <= row[i] {
			return true
		}
	}
	return false
}

func HiddenLeft(row []int, i int, j int) bool {
	value := row[j]
	for i := j - 1; i >= 0; i-- {
		if value <= row[i] {
			return true
		}
	}
	return false
}
