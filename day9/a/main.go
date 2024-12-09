package main

import (
	"fmt"
	"os"
)

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

	i = 0
	j := len(disk) - 1
	for i < j && disk[i] != '.' {
		i++
	}
	for j > i && disk[j] == '.' {
		j--
	}
	for i < j {

		disk[i], disk[j] = disk[j], disk[i]
		for i < j && disk[i] != '.' {
			i++
		}
		for j > i && disk[j] == '.' {
			j--
		}
	}

	sum := 0

	for i, value := range disk {
		if value == '.' {
			break
		}

		sum += int(value-'A') * i
	}
	fmt.Println(sum)
}
