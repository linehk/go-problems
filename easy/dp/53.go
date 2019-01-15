package dp

func maxSubArray(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	// sum, maxSum := math.MinInt32, math.MinInt32
	for _, v := range nums {
		sum = max(sum+v, v)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
