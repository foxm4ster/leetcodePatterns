package main

func maxProfit(prices []int) int {
	var profit, curr int

	curr = prices[0]
	for i := range prices {
		profit = max(profit, prices[i]-curr)
		curr = min(curr, prices[i])
	}

	return profit
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
