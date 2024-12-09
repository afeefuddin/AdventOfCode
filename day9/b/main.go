package main

import (
	"fmt"
	"os"
)

func captureFile(disk []rune, idx int) (int, int) {
	target := disk[idx]

	st := idx
	e := idx

	for st >= 0 && disk[st] == target {
		st--
	}

	for e < len(disk) && disk[e] == target {
		e++
	}

	return st + 1, e - 1
}

func main() {
	fi, err := os.ReadFile("../test.txt")
	if err != nil {
		panic(err)
	}
	input := string(fi)

	var disk []rune
	var i int = 0
	for idx, value := range input {
		temp := int(value - '0')
		char := '.'

		if idx%2 == 0 {
			char = rune('A' + i)
			i++
		}

		for j := 0; j < temp; j++ {
			disk = append(disk, char)
		}
	}

	j := len(disk) - 1

	for j >= 0 && disk[j] == '.' {
		j--
	}

	for j >= 0 {
		st, e := captureFile(disk, j)

		i := 0
		for i < j && disk[i] != '.' {
			i++
		}
		for i < st {
			fsStart, fsEnd := captureFile(disk, i)
			if fsEnd-fsStart >= e-st {
				lent := e - st
				for k := fsStart; k <= fsStart+lent; k++ {
					disk[k] = disk[st]
				}
				for k := st; k <= e; k++ {
					disk[k] = '.'
				}
				break
			} else {
				i = fsEnd + 1
			}

			for i < st && disk[i] != '.' {
				i++
			}
		}
		j = st - 1
	}

	sum := 0

	for i, value := range disk {
		if value == '.' {
			continue
		}

		sum += int(value-'A') * i
	}
	fmt.Println(sum)
}
