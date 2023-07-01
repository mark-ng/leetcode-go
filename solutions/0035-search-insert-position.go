// Day 18

func searchInsert(nums []int, target int) int {
    ask := func(i int) bool {
		return nums[i] >= target
	}
    return findFirst(ask, len(nums))
}

func findFirst(ask func(int) bool, n int) int {
    lo := 0
    hi := n - 1

    if n == 0 || ask(lo) == true {
        return lo 
    }
    if ask(hi) == false {
        return n
    }

    for hi - lo > 1 {
        mid := (hi + lo) / 2
        if ask(mid) == true {
            hi = mid
        } else {
            lo = mid
        }
    } 
    return hi
}

// REDO