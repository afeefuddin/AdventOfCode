// brute force approach
package main

import (
	"adventofcode/library"
	ll "adventofcode/linkedlist"
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

func blink(head *ll.Node[int]) {
	temp := head

	for temp != nil {
		value := temp.Val
		flag := false

		if value == 0 {
			temp.Val = 1
		} else if isEven(value) {
			a, b := splitNum(value)
			temp.Val = a
			ll.InsertForward(temp, b)
			flag = true
		} else {
			temp.Val = temp.Val * 2024
		}

		temp = temp.Next

		if flag && temp != nil {
			temp = temp.Next
		}
	}

}

func main() {
	fi, err := os.ReadFile("../test.txt")
	if err != nil {
		panic(err)
	}

	arr := library.ConvertStringArrayToInt(strings.Split(string(fi), " "))

	head := ll.MakeLL(arr)
	ll.PrintLL(head)
	for i := 0; i < 75; i++ {
		fmt.Printf("Loop %v\n", i)
		blink(head)
	}
	fmt.Println(ll.Size(head))
}
