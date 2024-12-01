package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int, st int, e int) int {
	pivot := arr[e]

	i, j := st, e

	for i < j {
		if arr[i] <= pivot {
			i++
		} else if arr[j] >= pivot {
			j--
		} else {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[e] = arr[e], arr[i]
	return i
}

func quickSort(arr []int, st int, e int) {
	if st < e {
		pi := partition(arr, st, e)

		quickSort(arr, st, pi-1)
		quickSort(arr, pi+1, e)
	}
}

func main() {
	fi, err := os.Open("../test.txt")

	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	var arr1 []int
	var arr2 []int

	for scanner.Scan() {
		line := scanner.Text()
		arrtemp := strings.Split(line, "   ")
		val1, err := strconv.Atoi(arrtemp[0])
		if err != nil {
			panic(err)
		}
		arr1 = append(arr1, val1)
		val2, err := strconv.Atoi(arrtemp[1])

		if err != nil {
			panic(err)
		}

		arr2 = append(arr2, val2)

	}



	quickSort(arr1, 0, len(arr1)-1)
	quickSort(arr2, 0, len(arr2)-1)

	var sum uint64 = 0


	for i := 0; i< len(arr1); i++ {
		sum += uint64(math.Abs(float64(arr1[i] - arr2[i])))
	}



	fmt.Printf("%v\n", sum)
}
