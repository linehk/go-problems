package dp

func rob(nums []int) int {
	evenSum, oddSum := 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			evenSum = max(evenSum+v, oddSum)
		} else {
			oddSum = max(evenSum, oddSum+v)
		}
	}
	return max(evenSum, oddSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
