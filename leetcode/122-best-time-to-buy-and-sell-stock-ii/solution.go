package leetcode

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] <= prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
