package main

import (
	"adventofcode/library"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculatePart1(target, value int, arr []int, idx int) bool {
	if idx == len(arr) {
		return target == value
	}
	ans := calculatePart1(target, value+arr[idx], arr, idx+1) || calculatePart1(target, value*arr[idx], arr, idx+1)
	return ans
}

func calculatePart2(target, value int, arr []int, idx int) bool {
	if idx == len(arr) {
		return target == value
	}

	ans := calculatePart2(target, value+arr[idx], arr, idx+1) || calculatePart2(target, value*arr[idx], arr, idx+1) ||
		calculatePart2(target, library.ConcatenateInt(value, arr[idx]), arr, idx+1)
	return ans
}

func main() {
	fi, err := os.Open("../test.txt")

	if err != nil {
		panic(err)
	}

	sum := 0
	sum2 := 0
	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		line := scanner.Text()

		splittedLine := strings.Split(line, ":")

		value, err := strconv.Atoi(splittedLine[0])

		if err != nil {
			panic(err)
		}

		arr := library.ConvertStringArrayToInt(strings.Split(strings.TrimSpace(splittedLine[1]), " "))

		if calculatePart1(value, arr[0], arr, 1) {
			sum += value
		}

		if calculatePart2(value, arr[0], arr, 1) {
			sum2 += value
		}
	}

	fmt.Println(sum, sum2)
}
