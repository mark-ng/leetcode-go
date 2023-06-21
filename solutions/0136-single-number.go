// Day 10  

func singleNumber(nums []int) int {
    ans := 0
    for _, v := range nums {
        ans = ans ^ v
    }
    return ans
}