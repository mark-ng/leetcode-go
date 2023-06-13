// Day 5

func removeDuplicates(nums []int) int {
    last := nums[0]
    count := 1
    end := len(nums)
    i := 1
    for i < end {
        if nums[i] == last {
            if count >= 2 {
                // mark as delete
                nums[i] = 100001
            } else {
                count++
            }  
        } else {
            count = 1
            last = nums[i]
        }
        i++
    } 
    return vaccum(nums)
}

func vaccum(nums []int) int {
    n := len(nums)
    for i := 0; i < len(nums); i++ {
        if i >= n {
            return n
        }
        if nums[i] == 100001 {
            n--
            for j := i; j < len(nums) - 1; j++ {
                nums[j] = nums[j + 1]
            }
            i--
        }
    }
    return n
}