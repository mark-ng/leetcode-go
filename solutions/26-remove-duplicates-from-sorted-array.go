// Day 1

func removeDuplicates(nums []int) int {
    before := -101
    k := 0
    for _, v := range nums {
        if v != before {
            nums[k] = v
            k++
            before = v
        }
    }
    return k
}