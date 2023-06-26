// Day 15

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := map[byte]int{}
    for i := 0; i < len(s); i++ {
        m[s[i]] += 1
    }
    for i := 0; i < len(t); i++ {
        m[t[i]] -= 1
        if m[t[i]] < 0 {
            return false
        }
    }
    return true
}