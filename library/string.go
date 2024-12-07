package library

import "strconv"

func ConvertStringArrayToInt(arr []string) []int {
	var ans []int
	for _, val := range arr {
		intValue, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		ans = append(ans, intValue)
	}
	return ans
}

func ConcatenateInt(a int, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	concatenated := aStr + bStr

	value, err := strconv.Atoi(concatenated)

	if err != nil {
		panic(err)
	}
	return value
}
