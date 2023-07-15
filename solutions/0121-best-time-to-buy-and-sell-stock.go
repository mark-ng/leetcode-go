// Day 23

func maxProfit(prices []int) int {
    n := len(prices)
    min := prices[0]
    maxProfit := 0
    for i := 1; i < n; i++ {
        if prices[i] - min > maxProfit {
            maxProfit = prices[i] - min
        }
        if prices[i] < min {
            min = prices[i]
        }
    }
    return maxProfit
    
}