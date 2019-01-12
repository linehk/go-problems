package array

func maxProfit(prices []int) int {
        profit := 0
        for i := 0; i < len(prices)-1; i++ {
                diff := prices[i+1]-prices[i]
                if diff > 0 {
                        profit += diff
                }
        }
        return profit
}