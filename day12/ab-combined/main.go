package main

import (
	"adventofcode/library"
	"adventofcode/queue"
	"fmt"
)

type coOrd struct {
	x int
	y int
}

func part1(data [][]rune, freq map[int]int, groups [][]int) int {

	ans := 0
	n := len(data)
	m := len(data[0])

	for i, row := range data {
		for j, value := range row {
			perim := 0
			if i == 0 {
				perim++
			} else {
				if data[i-1][j] != value {
					perim++
				}
			}

			if j == 0 {
				perim++
			} else {
				if data[i][j-1] != value {
					perim++
				}
			}

			if i == n-1 {
				perim++
			} else {
				if data[i+1][j] != value {
					perim++
				}
			}

			if j == m-1 {
				perim++
			} else {
				if data[i][j+1] != value {
					perim++
				}
			}
			ans += perim * freq[groups[i][j]]
		}
	}

	return ans

}

func bfs(groups [][]int, data [][]rune, pos coOrd, color int) {

	qu := queue.MakeQueue[coOrd]()

	qu.Push(pos)

	groups[pos.x][pos.y] = color
	for !qu.IsEmpty() {
		cur := qu.Pop()

		curVal := data[cur.x][cur.y]
		i, j := cur.x, cur.y

		if i > 0 && data[i-1][j] == curVal && groups[i-1][j] == 0 {
			qu.Push(coOrd{x: i - 1, y: j})
			groups[i-1][j] = color
		}
		if i < len(data)-1 && data[i+1][j] == curVal && groups[i+1][j] == 0 {
			qu.Push(coOrd{x: i + 1, y: j})
			groups[i+1][j] = color
		}
		if j > 0 && data[i][j-1] == curVal && groups[i][j-1] == 0 {
			qu.Push(coOrd{x: i, y: j - 1})
			groups[i][j-1] = color
		}
		if j < len(data[0])-1 && data[i][j+1] == curVal && groups[i][j+1] == 0 {
			qu.Push(coOrd{x: i, y: j + 1})
			groups[i][j+1] = color
		}
	}

}

func part2(groups [][]int, freq map[int]int) {
	n := len(groups)
	m := len(groups[0])
	ans := 0

	checkX := []int{-1, 1}
	checkY := []int{-1, 1}

	for key := range freq {
		corners := 0
		for i, rows := range groups {
			for j, val := range rows {
				if val != key {
					check1 := false
					check2 := false
					for _, x := range checkX {
						ni := i + x
						if ni < 0 || ni >= n {
							continue
						}

						check1 = (groups[ni][j] == key)
						for _, y := range checkY {
							nj := j + y
							if nj < 0 || nj >= m {
								continue
							}
							check2 = (groups[i][nj] == key)
							if check1 && check2 && groups[ni][nj] == key {
								corners++
							}
						}
					}

					continue
				}

				check1 := false
				check2 := false
				for _, x := range checkX {
					ni := i + x
					if ni < 0 || ni >= n {
						check1 = true
					} else {
						check1 = groups[ni][j] != key
					}
					for _, y := range checkY {
						nj := j + y
						if nj < 0 || nj >= m {
							check2 = true
						} else {
							check2 = groups[i][nj] != key
						}
						if check1 && check2 {
							corners++
						}
					}
				}

			}
		}
		ans += corners * freq[key]
	}

	fmt.Println("Soln 2: ", ans)
}

func main() {
	data := library.ReadGrid()

	n, m := len(data), len(data[0])

	groups := make([][]int, n)
	for i := 0; i < n; i++ {
		groups[i] = make([]int, m)
	}

	freq := make(map[int]int)

	groupVal := 1

	for i, row := range groups {
		for j, value := range row {
			if value == 0 {
				bfs(groups, data, coOrd{x: i, y: j}, groupVal)
				groupVal++
			}
		}
	}

	for _, row := range groups {
		for _, value := range row {
			freq[value]++
		}
	}

	// fmt.Printf("%v", groups)

	fmt.Println("Soln 1: ", part1(data, freq, groups))
	part2(groups, freq)
}
