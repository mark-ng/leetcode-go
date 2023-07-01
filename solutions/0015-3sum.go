// Day 16

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for _, j := range twoSumSort(nums[i+1:], -nums[i]) {
			tmp := append(j, nums[i])
			ans = append(ans, tmp)
		}
	}
	return ans
}

func twoSumSort(nums []int, offset int) [][]int {
	var ans [][]int
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum > offset {
			j--
		} else if sum < offset {
			i++
		} else {
			ans = append(ans, []int{nums[i], nums[j]})
			i++
			for nums[i] == nums[i-1] && i < j {
				i++
			}
		}
	}
	return ans
}

// REDO