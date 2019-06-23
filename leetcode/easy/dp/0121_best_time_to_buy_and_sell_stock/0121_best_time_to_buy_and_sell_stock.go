package dp

func maxProfit(prices []int) int {
	maxProfit := 0
	profit := 0
	for i := 1; i < len(prices); i++ {
		// profit 为累计利润
		profit += prices[i] - prices[i-1]
		// 表示根本没有利润，不交易，重置买入的时间
		if profit < 0 {
			profit = 0
		}
		// 计入 maxProfit
		maxProfit = max(maxProfit, profit)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
