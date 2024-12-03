package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() string {
	fi, err := os.ReadFile("../test.txt")

	if err != nil {
		panic(err)
	}
	content := string(fi)
	return content
}

func cleanData(data string) []string {
	re := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)|do\(\)|don't\(\)`)

	cleanedData := re.FindAllString(data, -1)
	return cleanedData
}

func main() {
	data := readFile()
	sum := 0
	enabled := true
	allData := cleanData(data)
	for _, val := range allData {
		if val == "do()" {
			enabled = true
			continue
		} else if val == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		substr := val[4:]
		substr = substr[:len(substr)-1]
		values := strings.Split(substr, ",")
		firstNumber, err := strconv.Atoi(values[0])

		if err != nil {
			panic(err)
		}

		secondNumber, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		sum += (firstNumber * secondNumber)
	}
	fmt.Printf("%v", sum)
}
