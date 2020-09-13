package dp

func maxSubArray(nums []int) int {
	// sum, maxSum := math.MinInt32, math.MinInt32
	// -1 = 11111111 11111111 11111111 11111111，左移 31 位后变成 10000000 00000000 00000000 00000000
	sum, maxSum := -1<<31, -1<<31
	for _, v := range nums {
		// sum 为 sum+v, v 中的最大值，第一步必定为 v
		// 当 sum+v > v 时，保留 sum
		// 当 v > sum+v 时，舍弃 sum，重新从 v 开始累积
		sum = max(sum+v, v)
		// maxSum 为 maxSum, sum 中的最大值，第一步必定为 sum，即 v
		// 当 sum 大于以往的累计和 maxSum 时，才会更新 maxSum 的值
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
