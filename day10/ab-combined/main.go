package main

import (
	"bufio"
	"fmt"
	"os"
)

type coOrd struct {
	x int
	y int
}

func search(x, y, target int, matrix [][]int, vis map[coOrd]bool) int {

	if target == 10 {
		vis[coOrd{x: x, y: y}] = true
		return 1
	}

	rows, cols := len(matrix), len(matrix[0])
	ans := 0
	if x > 0 && matrix[x-1][y] == target {
		ans += search(x-1, y, target+1, matrix, vis)
	}
	if x < rows-1 && matrix[x+1][y] == target {
		ans += search(x+1, y, target+1, matrix, vis)
	}
	if y > 0 && matrix[x][y-1] == target {
		ans += search(x, y-1, target+1, matrix, vis)
	}
	if y < cols-1 && matrix[x][y+1] == target {
		ans += search(x, y+1, target+1, matrix, vis)
	}
	return ans
}

func main() {

	fi, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}
	var matrix [][]int
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		var temp []int
		for _, val := range line {
			temp = append(temp, int(val-'0'))
		}
		matrix = append(matrix, temp)
	}

	ans := 0
	ans2 := 0

	for i, rows := range matrix {
		for j, value := range rows {
			if value == 0 {
				vis := make(map[coOrd]bool)
				ans2 += search(i, j, 1, matrix, vis)
				ans += len(vis)
			}
		}
	}

	fmt.Println(ans, ans2)

}
