// Day 1

func containsNearbyDuplicate(nums []int, k int) bool {
    m := map[int]int{}
    for i, v := range nums {
        if idx, ok := m[v]; ok {
            if i - idx <= k {
                return true
            } else {
                m[v] = i
            }
        } else {
            m[v] = i
        }
    }
    return false
}