package main

// Ugly code but i ain't refactoring it,
// problem made me cry already.

import (
	"adventofcode/library"
	"bufio"
	"fmt"
	"os"
)

func readFile() [][]rune {
	fi, err := os.Open("../test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fi)
	var matrix [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		temp := []rune(line)
		matrix = append(matrix, temp)
	}

	return matrix

}

func inRange(x, y int, maze [][]rune) bool {
	n := len(maze)    // number of rows
	m := len(maze[0]) // number of columng
	if x >= 0 && y >= 0 && x < m && y < n {
		return true
	}
	return false
}

// x is the column
// y is the row

func changeDirection(dx, dy, x, y int, maze [][]rune) (int, int) {
	possibleX := x + dx
	possibleY := y + dy

	if inRange(possibleX, possibleY, maze) {
		if possibleX == 8 && possibleY == 3 {

			fmt.Printf("%c",maze[possibleY][possibleX])
		}
		if maze[possibleY][possibleX] == '#' {
			if dy == 1 {
				return -1, 0
			} else if dy == -1 {
				return 1, 0
			} else if dx == 1 {
				return 0, 1
			} else {
				return 0, -1
			}
		}
	}

	return dx, dy
}

func findStartingPosition(matrix [][]rune) (int, int) {
	for i, row := range matrix {
		for j, col := range row {
			if col == '^' {
				return i, j
			}
		}
	}

	panic("Character not found")
}

func possilbePositions(data [][]rune) ([][]rune, int) {

	// y is row = i
	// x is col = j

	width := len(data[0])
	height := len(data)
	y, x := findStartingPosition(data)
	dy, dx := -1, 0
	shouldEnd := false

	ans := make([][]rune, len(data))
	for i := range data {
		ans[i] = make([]rune, len(data[i]))
		copy(ans[i], data[i])
	}

	count := 0

	for {
		if ans[y][x] != 'X' {
			ans[y][x] = 'X'
			count++
		}

		if shouldEnd {
			break
		}

		dx, dy = changeDirection(dx, dy, x, y, data)

		x += dx
		y += dy

		if x < 0 || x >= width {
			break
		}
		if y < 0 || y >= height {
			break
		}

	}
	return ans, count
}

type coordinates struct {
	x int
	y int
}

func isLoop(data [][]rune, y, x int) bool {
	dy, dx := -1, 0
	width := len(data[0])
	height := len(data)
	shouldEnd := false

	// 0 --> for up
	// 1 --> for down
	// 2 --> for right
	// 3 --> for left
	coord := make(map[coordinates][4]int)
	

	for {

		if value, found := coord[coordinates{x: x, y: y}]; found {
			if dy == 1 {
				if value[1] == 1 {
					return true
				} else {
					value[1] = 1
				}
			} else if dy == -1 {
				if value[0] == 1 {
					return true
				} else {
					value[0] = 1
				}
			} else if dx == 1 {
				if value[2] == 1 {
					return true
				} else {
					value[2] = 1
				}
			} else {
				if value[3] == 1 {
					return true
				} else {
					value[3] = 1
				}
			}

		} else {
			coord[coordinates{x: x, y: y}] = [4]int{library.BoolToInt(-1*dy == 1),
				library.BoolToInt(1*dy == 1), library.BoolToInt(1*dx == 1), library.BoolToInt(-1*dx == 1)}
		}

		newDx, newDy := changeDirection(dx, dy, x, y, data)

		for {
			tempDx, tempDy := changeDirection(newDx, newDy, x, y, data)

			if tempDx == newDx && tempDy == newDy {
				break
			}
			newDx, newDy = tempDx, tempDy
		}

		if shouldEnd && newDx == dx && newDy == dy {
			break
		}

		dx, dy = newDx, newDy

		x += dx
		y += dy

		if x < 0 || x > width {
			shouldEnd = true
		}
		if y < 0 || y > height {
			shouldEnd = true
		}

	}

	return false

}

func calculatePossibleObstructions(maze [][]rune, positions [][]rune, y, x int) int {
	count := 0

	for i, row := range positions {
		for j, col := range row {
			if col == 'X' {
				maze[i][j] = '#'
				if i == 3 && j == 8 {
					fmt.Printf("yes loop check is  at %v %v \n", i, j)
				}
				if isLoop(maze, y, x) {
					fmt.Printf("loop at %v %v \n", i, j)
					count++
				}
				if y == i && x == j {
					maze[i][j] = '^'
				} else {
					maze[i][j] = '.'
				}
			}
		}
	}

	return count
}

func main() {
	maze := readFile()
	pos, count := possilbePositions(maze)
	y, x := findStartingPosition(maze)
	obs := calculatePossibleObstructions(maze, pos, y, x)
	fmt.Println(count, obs)
}
