// Day 9

func coinChange(coins []int, amount int) int {
	memo := map[int]int{}
	return coinChangeMemo(coins, amount, memo)
}

func coinChangeMemo(coins []int, amount int, memo map[int]int) int {
	val, ok := memo[amount]
	if ok {
		return val
	}
	if amount < 0 {
        return -1
    }
    if amount == 0 {
        return 0
    }
    var minCoinChange int = -1
    for _, coin := range coins {
        r := coinChangeMemo(coins, amount - coin, memo)
        if minCoinChange == -1 && r != -1 {
            minCoinChange = r
        } else if minCoinChange != -1 && r != -1 {
            minCoinChange = min(minCoinChange, r)
        }
    }
    if minCoinChange == -1 {
		memo[amount] = -1
        return -1
    }
	memo[amount] = minCoinChange + 1 
    return minCoinChange + 1	
}


func min(a int, b int) int {
    if a <= b {
        return a
    }
    return b
}