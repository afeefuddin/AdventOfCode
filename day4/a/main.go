package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() [][]rune {
	fi, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	var matrix [][]rune

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		temp := []rune(line)
		matrix = append(matrix, temp)
		i++
	}
	return matrix
}

func searchMAS(a, b, c rune) bool {
	return a == 'M' && b == 'A' && c == 'S'
}

func searchXmas(matrix [][]rune, i, j int) int {
	count := 0

	n := len(matrix)
	m := len(matrix[0])

	// search Top

	if i >= 3 {
		if searchMAS(matrix[i-1][j], matrix[i-2][j], matrix[i-3][j]) {
			count++
		}
	}

	// serach bottom

	if i < n-3 {
		if searchMAS(matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]) {
			count++
		}
	}

	// search right

	if j < m-3 {
		if searchMAS(matrix[i][j+1], matrix[i][j+2], matrix[i][j+3]) {
			count++
		}

	}

	// search left
	if j >= 3 {
		if searchMAS(matrix[i][j-1], matrix[i][j-2], matrix[i][j-3]) {
			count++
		}
	}
	// search top-left

	if i >= 3 && j >= 3 {
		if searchMAS(matrix[i-1][j-1], matrix[i-2][j-2], matrix[i-3][j-3]) {
			count++
		}
	}

	// search top-right

	if i >= 3 && j < m-3 {
		if searchMAS(matrix[i-1][j+1], matrix[i-2][j+2], matrix[i-3][j+3]) {
			count++
		}
	}

	// search bottom-left
	if i < n-3 && j >= 3 {
		if searchMAS(matrix[i+1][j-1], matrix[i+2][j-2], matrix[i+3][j-3]) {
			count++
		}
	}

	// search bottom-right
	if i < n-3 && j < m-3 {
		if searchMAS(matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3]) {
			count++
		}
	}

	return count
}

func main() {
	var matrix [][]rune = readFile()
	sum := 0
	for i, row := range matrix {
		for j, value := range row {
			if value == 'X' {
				sum += searchXmas(matrix, i, j)
			}
		}
	}
	fmt.Println(sum)

}
