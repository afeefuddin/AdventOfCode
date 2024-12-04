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

func searchXmas(matrix [][]rune, i, j int) int {

	n := len(matrix)
	m := len(matrix[0])

	if i >= 1 && j >= 1 && i < n-1 && j < m-1 {
		// diag 1 check
		pass := 0

		if (matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S') || (matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M') {
			pass++
		}
		if (matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S') || (matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M') {
			pass++
		}
		if pass == 2 {
			return 1
		}
	}
	return 0
}

func main() {
	var matrix [][]rune = readFile()
	sum := 0
	for i, row := range matrix {
		for j, value := range row {
			if value == 'A' {
				sum += searchXmas(matrix, i, j)
			}
		}
	}
	fmt.Println(sum)
}
