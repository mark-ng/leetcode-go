// Day 7

func longestCommonPrefix(strs []string) string {
    ans := strs[0]
    for i := 1; i < len(strs); i++ {
        ans = commonPrefix(ans, strs[i])
    }
    return ans
}

func commonPrefix(s1 string, s2 string) string {
    minStr := min(s1, s2)
    minLen := len(minStr)
    i := 0
    for i < minLen {
        if s1[i] != s2[i] {
            return minStr[:i]
        }
        i++
    }
    return minStr[:i]
}

func min(a, b string) string {
    if len(a) < len(b) {
        return a
    }
    return b
}