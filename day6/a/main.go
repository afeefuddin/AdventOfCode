package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

func findStartingPosition(matrix [][]rune) (int, int) {
	for i, row := range matrix {
		for j, col := range row {
			if col == '^' {
				// terminal have 1 based indexing
				return i + 1, j + 1
			}
		}
	}

	panic("Character not found")
}

func printMaze(maze [][]rune) {
	for _, row := range maze {
		for _, col := range row {
			switch col {
			case '#':
				fmt.Printf("\033[31m%c\033[0m", col)
			case '.':
				fmt.Printf("\033[37m%c\033[0m", col)
			case 'X':
				fmt.Printf("\033[33m%c\033[0m", col)
			case '^':
				fmt.Printf("\033[37m%c\033[0m", '.')
			default:
				fmt.Printf("%c", col)
			}
		}
		fmt.Println()
	}
}

func printCharacter(dy, dx, x, y int) {
	if dy == 1 {
		fmt.Printf("\033[%d;%dH\033[34mv\033[0m", y, x)
	} else if dy == -1 {
		fmt.Printf("\033[%d;%dH\033[34m^\033[0m", y, x)
	} else if dx == 1 {
		fmt.Printf("\033[%d;%dH\033[34m>\033[0m", y, x)
	} else {
		fmt.Printf("\033[%d;%dH\033[34m<\033[0m", y, x)
	}
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
	possibleX := x + dx - 1
	possibleY := y + dy - 1

	if inRange(possibleX, possibleY, maze) {
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

func main() {

	data := readFile()

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
		// Clear the screen
		fmt.Print("\033[H\033[J")
		// Clear the cursor
		fmt.Print("\033[?25l")

		printMaze(data)
		printCharacter(dy, dx, x, y)

		if ans[y-1][x-1] != 'X' {
			ans[y-1][x-1] = 'X'
			count++
		}

		if shouldEnd {
			break
		}

		dx, dy = changeDirection(dx, dy, x, y, data)

		x += dx
		y += dy

		if x <= 1 || x >= width {
			shouldEnd = true
		}
		if y <= 1 || y >= height {
			shouldEnd = true
		}

		time.Sleep(200 * time.Millisecond)
	}
	fmt.Print("\033[H\033[J")
	fmt.Print("\033[?25h")
	printMaze(ans)
	fmt.Printf("\nYour character covered %v blocks\n", count)
}
