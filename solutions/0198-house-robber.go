// Day 22

func rob(nums []int) int {
    // Base case of sub problem
    // if len(nums) == 1, can only steal that
    // if len(nums) == 2, can only steal the larger one
    // if len(nums) >= 3, max(ans(i - 1) , ans(i - 2) + nums[i])
    
    n := len(nums)
    ans := make([]int, n)
    ans[0] = nums[0]
    max := ans[0]

    if n > 1 {
        ans[1] = mmax(nums[0], nums[1])
        if ans[1] > max {
            max = ans[1]
        }
        for i := 2; i < n; i++ {
            adjacentSolution := ans[i - 1]
            friendSolution := ans[i - 2]

            ans[i] = mmax(adjacentSolution, friendSolution + nums[i])
            if ans[i] > max {
                max = ans[i]
            }
        }
    }
    return max   
}

func mmax(a int, b int) int {
    if a > b {
        return a
    }
    return b
}