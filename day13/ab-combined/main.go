package main

import (
	"adventofcode/library"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getNumbers(str string) (int, int) {
	re := regexp.MustCompile(`\d+`)
	match := library.ConvertStringArrayToInt(re.FindAllString(str, 2))

	return match[0], match[1]
}

func canConvertToInt(f float64) bool {
	return f == float64(int(f))
}

func parseEval(eval string, higherTarget bool) int {

	temp := strings.Split(eval, "\n")
	line1 := temp[0]
	line2 := temp[1]
	line3 := temp[2]
	aX, aY := getNumbers(line1)
	bX, bY := getNumbers(line2)
	targetX, targetY := getNumbers(line3)

	if higherTarget {
		targetX += 10000000000000
		targetY += 10000000000000
	}

	// aX*m + bX*n = targetX
	// aY*m + bY*n = targetY

	// m = targetX - bX*n / aX
	// (aY/aX)*(targetX - bX*n) + by*n = targetY
	// -aY*bX * n/aX + bY*n = targetY - ay*targetX/aX
	// n* (bY*aX - aY*bX)/aX = (targetY*aX - ay*targetX)/aX

	var n float64 = (float64(targetY)*float64(aX) - float64(aY)*float64(targetX)) / (float64(bY)*float64(aX) - float64(aY)*float64(bX))
	fmt.Println(n)
	if !canConvertToInt(n) {
		return 0
	}

	m := (float64(targetX) - float64(bX)*n) / float64(aX)
	if !canConvertToInt(m) {
		return 0
	}

	return 3*int(m) + int(n)
}

func main() {
	fi, err := os.ReadFile("../test.txt")
	if err != nil {
		panic(err)
	}

	tokens := 0
	allTokens := 0

	tcases := strings.Split(string(fi), "\n\n")
	for _, tc := range tcases {
		tokens += parseEval(tc, false)
		allTokens += parseEval(tc, true)
	}
	fmt.Println(tokens, allTokens)

}
