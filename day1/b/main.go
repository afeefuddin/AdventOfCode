package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fi, err := os.Open("../test.txt")

	if err != nil {
		panic(err)
	}

	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	var arr1 []int
	frequency := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()

		cleanLine := strings.Join(strings.Fields(line), " ")

		arrTemp := strings.Split(cleanLine, " ")

		val1, err := strconv.Atoi(arrTemp[0])

		if err != nil {
			panic(err)
		}
		
		arr1 = append(arr1, val1)

		val2, err := strconv.Atoi(arrTemp[1])

		if err != nil {
			panic(err)
		}

		_, exists := frequency[val2]

		if !exists {
			frequency[val2] = 0
		}
		frequency[val2] = frequency[val2] + 1
	}

	var ans uint64 = 0

	for i :=0 ; i< len(arr1); i++ {
		ans += uint64(arr1[i]) * uint64(frequency[arr1[i]])
	}

	fmt.Println(ans)
}
