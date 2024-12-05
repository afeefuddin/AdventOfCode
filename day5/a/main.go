package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"adventofcode/library"
)

func readFile() []string {
	fi, err := os.ReadFile("../test.txt")
	if err != nil {
		panic(err)
	}
	content := string(fi)

	data := strings.Split(content, "\n\n")


	if len(data) != 2 {
		panic("Invalid data")
	}
	return data
}

func main() {
	data := readFile()
	relation, orders := data[0], data[1]

	myMap := make(map[int][]int)
	scanner := bufio.NewScanner(strings.NewReader(relation))

	for scanner.Scan() {
		relate := strings.Split(scanner.Text(), "|")
		a, err := strconv.Atoi(relate[0])

		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(relate[1])

		if err != nil {
			panic(err)
		}

		myMap[a] = append(myMap[a], b)
	}

	scanner = bufio.NewScanner(strings.NewReader(orders))

	sum := 0

	for scanner.Scan() {
		order := library.ConvertStringArrayToInt(strings.Split(scanner.Text(), ","))
		valid := true
		for i:=0;i<len(order);i++ {
			for j:=i+1; j<len(order); j++ {
				if library.Contains(myMap[order[j]], order[i]) {
					valid = false
					break
				}
			}
			if !valid {
				break;
			}
		}
		if valid {
			sum += order[len(order)/2]
		}
	}
	fmt.Println(sum)
}
