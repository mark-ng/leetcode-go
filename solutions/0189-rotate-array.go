// Day 6

func rotate(nums []int, k int)  {
    clone := make([]int, len(nums))
    copy(clone, nums)
    for i := 0; i < len(nums); i++ {
        clone[(i + k) % len(nums)] = nums[i]
    }
    copy(nums, clone)
}