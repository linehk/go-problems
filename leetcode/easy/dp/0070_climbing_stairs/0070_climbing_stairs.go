package dp

func climbStairs(n int) int {
	// n < 2 时，只有 1 种方法
	if n < 2 {
		return 1
	}
	// 斐波那契数列
	rec := make([]int, n+1)
	rec[0] = 1 // 0 阶时只有 1 种方法
	rec[1] = 1 // 1 阶时只有 1 种方法
	for i := 2; i <= n; i++ {
		// i = 2 时，rec[2] = rec[1] + rec[0] = 2，即有 2 种方法
		rec[i] = rec[i-1] + rec[i-2]
	}
	return rec[n]
}
