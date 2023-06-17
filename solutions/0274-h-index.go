// Day 8

func hIndex(citations []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(citations)))
    n := len(citations)
    ans := 0
    for i := n; i > 0; i-- {
        count := 0
        for _, citation := range citations {
            if citation >= i {
                count++
            }
        }
        if count >= i {
            return i   
        }
    }
    return ans
}