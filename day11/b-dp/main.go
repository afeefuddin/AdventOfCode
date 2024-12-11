// dynamic programming approach
package main

import (
	"adventofcode/library"
	"fmt"
	"os"
	"strings"
)

func isEven(value int) bool {
	dig := library.NumberOfDigits(value)
	return dig%2 == 0
}

func splitNum(value int) (int, int) {
	dig := library.NumberOfDigits(value)
	b := 0
	a := 0

	i := 1
	for j := 0; j < dig/2; j++ {
		lastDig := value % 10
		b += lastDig * i
		i = i * 10
		value = value / 10
	}
	i = 1
	for j := 0; j < dig/2; j++ {
		lastDig := value % 10
		a += lastDig * i
		i = i * 10
		value = value / 10
	}

	return a, b
}

func blinkAStone(value, times int, dp map[int]map[int]int) int {
	if times == 0 {
		return 1
	}
	if mapVal, found := dp[value]; found {
		if cachedResult, exists := mapVal[times]; exists {
			return cachedResult
		}
	}
	if dp[value] == nil {
		dp[value] = make(map[int]int)
	}

	var newVal []int

	if value == 0 {
		newVal = append(newVal, 1)
	} else if isEven(value) {
		a, b := splitNum(value)
		newVal = append(newVal, a, b)
	} else {
		newVal = append(newVal, value*2024)
	}

	sum := 0
	for _, val := range newVal {
		sum += blinkAStone(val, times-1, dp)
	}
	dp[value][times] = sum

	return sum
}

func main() {
	fi, err := os.ReadFile("../test.txt")
	if err != nil {
		panic(err)
	}

	arr := library.ConvertStringArrayToInt(strings.Split(string(fi), " "))

	dp := make(map[int]map[int]int)

	ans := 0
	for _, val := range arr {
		ans += blinkAStone(val, 75, dp)
	}
	fmt.Println(ans)
}
