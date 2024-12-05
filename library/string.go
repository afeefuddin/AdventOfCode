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