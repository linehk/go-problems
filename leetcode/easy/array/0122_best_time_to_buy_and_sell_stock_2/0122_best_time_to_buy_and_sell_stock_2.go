package array

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		// diff 表示收益
		diff := prices[i+1] - prices[i]
		// 收益大于 0 则表示可以进行一次买卖
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}
