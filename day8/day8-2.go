package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var verbose bool

func main() {
	flag.BoolVar(&verbose, "v", false, "should verbose?")
	flag.Parse()

	file, err := os.Open("./input")

	if err != nil {
		log.Fatal("Can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := ReadAndConvertToMatrix(scanner, 99, 99)

	start := time.Now()

	highest, err := HighestScenicScore(matrix)

	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(start)

	fmt.Println("highest:", highest)
	fmt.Println("duration ->", duration.String())
}

func HighestScenicScore(matrix [][]int) (int, error) {
	highest := 0
	row := 1
	for row < len(matrix)-1 {
		col := 1
		for col < len(matrix[row])-1 {
			current := matrix[row][col]

			top, bottom, left, right, err := GetDirections(matrix, col, row)

			if err != nil {
				return 0, err
			}

			topScore, bottomScore, leftScore, rightScore := CalculateDirectionScores(current, top, bottom, left, right)

			totalScore := topScore * bottomScore * leftScore * rightScore

			Debug(fmt.Sprintf("#%v:%v = %v -> %v - %v - %v - %v = %v\n", row, col, current, topScore, bottomScore, leftScore, rightScore, totalScore))

			if totalScore > highest {
				highest = totalScore
			}

			col++
		}
		row++
	}
	return highest, nil
}

// reads the file and converts its lines into a matrix
func ReadAndConvertToMatrix(scanner *bufio.Scanner, rows, cols int) [][]int {
	matrix := NewMatrix(rows, cols)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		matrix[row] = ToNumberVec(line)
		row++
	}
	return matrix
}

// transforms a string into a vector of numbers
func ToNumberVec(str string) []int {
	vec := make([]int, len(str))
	for idx, char := range str {
		a := string(char)
		value, _ := strconv.Atoi(a)
		vec[idx] = value
	}
	return vec
}

// initialize a matrix with the given sizes
func NewMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func GetCol(matrix [][]int, colIdx, from, to int) ([]int, error) {
	if from < 0 || from > len(matrix)-1 || to <= from || to > len(matrix) {
		return nil, fmt.Errorf("col out of bound")
	}
	length := to - from
	col := make([]int, length)
	for rows := 0; rows < length; rows++ {
		col[rows] = matrix[rows+from][colIdx]
	}
	return col, nil
}

func GetRow(matrix [][]int, rowIdx, from, to int) ([]int, error) {
	if from < 0 || from > len(matrix[rowIdx])-1 || to <= from || to > len(matrix[rowIdx]) {
		return nil, fmt.Errorf("row out of bound")
	}
	return matrix[rowIdx][from:to], nil
}

func GetDirections(matrix [][]int, col, row int) (top, bottom, left, right []int, err error) {
	top, err = GetCol(matrix, col, 0, row)
	if err != nil {
		return top, bottom, left, right, fmt.Errorf("error while getting top vector")
	}
	bottom, err = GetCol(matrix, col, row+1, len(matrix))
	if err != nil {
		return top, bottom, left, right, fmt.Errorf("error while getting bottom vector")
	}
	left, err = GetRow(matrix, row, 0, col)
	if err != nil {
		return top, bottom, left, right, fmt.Errorf("error while getting left vector")
	}
	right, err = GetRow(matrix, row, col+1, len(matrix[col]))
	if err != nil {
		return top, bottom, left, right, fmt.Errorf("error while getting right vector")
	}
	return
}

// Calculates the number of trees visible in a given vector `vec` by a tree with
// height `targetHeight`
//
// Effectively, this function returns how many consecutive items in the list are
// less than or equal to targetHeight. It will starts from 0 to len(vec) if reverse
// is false and from len(vec) to 0 if reverse is true
//
// Returns:
// The number of visible trees
func CalculateTreeQuant(vec []int, targetHeight int, reverse bool) int {
	sum := 0
	if reverse {
		for i := len(vec) - 1; i >= 0; i-- {
			sum++
			if vec[i] >= targetHeight {
				break
			}
		}
	} else {
		for i := 0; i < len(vec); i++ {
			sum++
			if vec[i] >= targetHeight {
				break
			}
		}
	}
	return sum
}

func CalculateDirectionScores(current int, top, bottom, left, right []int) (topScore, bottomScore, leftScore, rightScore int) {
	topScore = CalculateTreeQuant(top, current, true)
	bottomScore = CalculateTreeQuant(bottom, current, false)
	leftScore = CalculateTreeQuant(left, current, true)
	rightScore = CalculateTreeQuant(right, current, false)
	return
}

func Debug(msg string) {
	if verbose {
		fmt.Print(msg)
	}
}
