// Day 2

func twoSum(nums []int, target int) []int {
    // k: offset, v: index
    m := map[int]int{}
    for i, v := range nums {
        if v, ok := m[v]; ok {
            return []int{i, v} 
        }
        m[target - v] = i
    }
    return nil
}