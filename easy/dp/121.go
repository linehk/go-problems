package dp

func maxProfit(prices []int) int {
	profit := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if temp > profit {
			profit = temp
		}
	}
	return profit
}
