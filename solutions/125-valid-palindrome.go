func isAlphanumeric(c byte) bool {
    return (c >= byte('A') && c <= byte('Z')) || (c >= byte('a') && c <= byte('z')) || (c >= byte('0') && c <= byte('9'))
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    i, j := 0, len(s) - 1
    for i < j {
        if !isAlphanumeric(s[i]) {
            i++
            continue
        }
        if !isAlphanumeric(s[j]) {
            j--
            continue
        }
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}