// Day 14

func wordPattern(pattern string, s string) bool {
    m := map[byte]string{}
    m2 := map[string]byte{}
    sArr := strings.Split(s, " ")
    if len(pattern) != len(sArr) {
        return false
    }
    for i := 0; i < len(pattern); i++ {
        if v, ok := m[pattern[i]]; ok {
            if v != sArr[i] {
                return false
            }
        } else {
           m[pattern[i]] = sArr[i]
        }
        if v, ok := m2[sArr[i]]; ok {
            if v != pattern[i] {
                return false
            }
        } else {
            m2[sArr[i]] = pattern[i]
        }
    }
    return true
}