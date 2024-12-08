package main

import (
	"adventofcode/library"
	"fmt"
)

type coOrd struct {
	x int
	y int
}

var myMap map[rune][]coOrd

func storeAntiNodes(x, y int, antiNodes [][]int) bool {
	rows, cols := len(antiNodes), len(antiNodes[0])
	if x < 0 || x >= rows || y < 0 || y >= cols {
		return false
	}
	antiNodes[x][y] = 1
	return true
}

func findAntiNodes(antiNodes [][]int) int {
	count := 0
	for _, value := range myMap {
		for i := 0; i < len(value); i++ {
			for j := i + 1; j < len(value); j++ {
				// xDiff :=
				a := value[i]
				b := value[j]

				if a.x > b.x {
					xDiff := a.x - b.x
					if a.y > b.y {
						yDiff := a.y - b.y
						i := 1
						for storeAntiNodes(b.x-i*xDiff, b.y-i*yDiff, antiNodes) {
							i++
						}
						i = 1
						for storeAntiNodes(a.x+i*xDiff, a.y+i*yDiff, antiNodes) {
							i++
						}
					} else {
						yDiff := b.y - a.y
						i := 1
						for storeAntiNodes(a.x+i*xDiff, a.y-i*yDiff, antiNodes) {
							i++
						}
						i = 1
						for storeAntiNodes(b.x-i*xDiff, b.y+i*yDiff, antiNodes) {
							i++
						}
					}
				} else {
					xDiff := b.x - a.x
					if b.y > a.y {
						yDiff := b.y - a.y
						i := 0
						for storeAntiNodes(a.x-i*xDiff, a.y-i*yDiff, antiNodes) {
							i++
						}
						i = 0
						for storeAntiNodes(b.x+i*xDiff, b.y+i*yDiff, antiNodes) {
							i++
						}
					} else {
						yDiff := a.y - b.y
						i := 0
						for storeAntiNodes(b.x+i*xDiff, b.y-i*yDiff, antiNodes) {
							i++
						}
						i = 0
						for storeAntiNodes(a.x-i*xDiff, a.y+i*yDiff, antiNodes) {
							i++
						}
					}
				}

			}
		}
	}
	for _, rows := range antiNodes {
		for _, value := range rows {
			if value == 1 {
				count++
			}
		}
	}
	return count
}

func findNodes(data [][]rune, antiNodes [][]int) int {
	myMap = make(map[rune][]coOrd)

	for i, row := range data {
		for j, value := range row {
			if value != '.' {
				if val, found := myMap[value]; found {
					// Append the new coordinate and update the map
					myMap[value] = append(val, coOrd{x: i, y: j})
				} else {
					// Initialize the slice with the new coordinate
					myMap[value] = []coOrd{{x: i, y: j}}
				}
			}
		}
	}

	return findAntiNodes(antiNodes)
}

func main() {
	data := library.ReadGrid()
	rows, cols := len(data), len(data[0])
	antiNodes := library.Make2dArray(rows, cols)
	fmt.Println(findNodes(data, antiNodes))
}
