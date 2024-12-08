package library

func Contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func Make2dArray(rows, cols int) [][]int {
	ans := make([][]int, rows)

	for i := range ans {
		ans[i] = make([]int, cols)
	}
	return ans
}
