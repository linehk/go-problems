package dp

func rob(nums []int) int {
	evenSum, oddSum := 0, 0
	for i, v := range nums {
		// 对于数组中的每个数，选择偶数和和奇数和中最大的那个
		if i%2 == 0 {
			// 偶数和
			evenSum = max(evenSum+v, oddSum)
		} else {
			// 奇数和
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
