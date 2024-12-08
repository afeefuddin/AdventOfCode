package library

import (
	"bufio"
	"os"
)

func ReadGrid() [][]rune {
	fi, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}

	var ans [][]rune
	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		line := scanner.Text()

		ans = append(ans, []rune(line))
	}
	return ans
}
