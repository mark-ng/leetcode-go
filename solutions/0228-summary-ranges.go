// Day 18

func summaryRanges(nums []int) []string {
    var ans []string
    if len(nums) == 0 {
        return ans
    }
    start := nums[0]
    end := nums[0] + 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] + 1 {
            continue
        } else {
            end = nums[i - 1] 
            if start == end {
                ans = append(ans, strconv.Itoa(start))
            } else {
                ans = append(ans, strconv.Itoa(start) + "->" + strconv.Itoa(end))
            }
            start = nums[i]
        }
    }
    end = nums[len(nums) - 1]
    if start == end {
        ans = append(ans, strconv.Itoa(start))
    } else {
        ans = append(ans, strconv.Itoa(start) + "->" + strconv.Itoa(end))
    }
    return ans 
}