// Day 23

func maxProfit(prices []int) int {
    n := len(prices)
    min := prices[0]
    profit := 0
    for i := 1; i < n;i++ {
        if prices[i] < prices[i - 1] {
            min = prices[i]
        } else {
            profit += prices[i] - min
            min = prices[i]
        }
    }
    return profit
}