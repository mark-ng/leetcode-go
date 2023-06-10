// Day 3

func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    sp := 0
    for _, v := range t {
        if sp >= len(s) {
            return true
        }
        if v == rune(s[sp]) {
            sp++
        }
    }
    return sp == len(s)
}