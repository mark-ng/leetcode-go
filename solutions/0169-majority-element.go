// Day 3

func majorityElement(nums []int) int {
    n := len(nums)
    m := map[int]int{}
    for _, v := range nums {
        m[v]++
    }
    for k, v := range m {
        if v > n / 2 {
            return k
        }
    }
    return 0
}