package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func increasingCheck(arr []int, drops int, i int) bool {
	if drops > 1 {
		return false
	}
	if i > len(arr)-2 {
		return true
	}

	// 5 1 2 3 4 5
	// 1 2 3 4 5 2

	diff := arr[i+1] - arr[i]

	if 1 <= diff && diff <= 3 {
		return increasingCheck(arr, drops, i+1)
	} else {
		tempArr1 := append([]int{}, arr[:i]...)
		tempArr1 = append(tempArr1, arr[i+1:]...)

		tempArr2 := append([]int{}, arr[:i+1]...)
		tempArr2 = append(tempArr2, arr[i+2:]...)
		firstCheck := increasingCheck(tempArr1, drops+1, 0)
		secondCheck := increasingCheck(tempArr2, drops+1, 0)
		return firstCheck || secondCheck
	}

}

func decreasingCheck(arr []int, drops int, i int) bool {
	if drops > 1 {
		return false
	}
	if i > len(arr)-2 {
		return true
	}
	// 5 1 2 3 4 5
	// 1 2 3 4 5 2

	diff := arr[i] - arr[i+1]

	if 1 <= diff && diff <= 3 {
		return decreasingCheck(arr, drops, i+1)
	} else {
		tempArr1 := append([]int{}, arr[:i]...)
		tempArr1 = append(tempArr1, arr[i+1:]...)

		tempArr2 := append([]int{}, arr[:i+1]...)
		tempArr2 = append(tempArr2, arr[i+2:]...)
		firstCheck := decreasingCheck(tempArr1, drops+1, 0)
		secondCheck := decreasingCheck(tempArr2, drops+1, 0)
		return firstCheck || secondCheck
	}
}

func isSafe(arr []int) bool {

	return increasingCheck(arr, 0, 0) || decreasingCheck(arr, 0,0)
}

func stringToInt(arr []string) []int {
	var ans []int
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
