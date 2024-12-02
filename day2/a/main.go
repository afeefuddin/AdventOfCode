package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func increasingCheck(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		} else if arr[i]-arr[i-1] > 3 {
			return false
		} 
	}
	return true
}

func decreasingCheck(arr []int) bool {
	for i := 0; i < len(arr) -1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		} else if arr[i]-arr[i+1] > 3 {
			return false
		}
	}
	return true
}

func isSafe(arr []int) bool {
	if arr[0] == arr[1] {
		return false
	}
	incr := arr[0] < arr[1]

	if incr {
		return increasingCheck(arr)
	} else {
		return decreasingCheck(arr)
	}
}

func stringToInt (arr []string) []int {
	var ans []int;
	for _, i := range arr {
		j, err := strconv.Atoi(i)

		if err != nil {
			panic(err)
		}
		ans = append(ans, j)
	}
	return ans
}

func main() {
	fi, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		arrtemp := strings.Split(line, " ")
		if len(arrtemp) <= 2 {
			safeCount++
			continue
		}
		if isSafe(stringToInt(arrtemp)) {
			safeCount++
		}		

	}
	fmt.Println(safeCount)

}
