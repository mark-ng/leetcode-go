// Day 2

func permute(nums []int) [][]int {
	result := [][]int{}
	dfs(nums, &[]int{}, &result)
	return result
}

func dfs(nums []int, accu *[]int, result *[][]int) {
	if allVisited(nums) {
		*result = append(*result, append([]int{}, *accu...))
		return
	}
	for i, num := range nums {
		if num == 11 {
			continue
		}
		nums[i]	= 11
		*accu = append(*accu, num)
		dfs(nums, accu, result)
		nums[i] = num
		*accu = (*accu)[:len(*accu) - 1]
	}
}

func allVisited(nums []int) bool {
	for _, num := range nums {
		// 11 is a hack~
		if num != 11 {
			return false
		}
	}
	return true
}