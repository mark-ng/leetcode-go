// Day 1

func merge(nums1 []int, m int, nums2 []int, n int)  {
    k, m, n := m + n - 1, m - 1, n - 1

    for m >= 0 || n >= 0 {
        if m < 0 && n >= 0 {
            nums1[k] = nums2[n]
            n--
        } else if m >= 0 && n < 0 {
            break
        } else if nums1[m] >= nums2[n] {
            nums1[k] = nums1[m]
            m--
        } else {
            nums1[k] = nums2[n]
            n--
        }
        k--
    }
}