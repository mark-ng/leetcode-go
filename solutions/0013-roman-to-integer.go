// Day 5

func romanToInt(s string) int {
    sl := len(s) - 1
    m := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    sum := 0
    last := 0
    for sl >= 0 {
        if m[s[sl]] >= last {
            sum += m[s[sl]]
            last = m[s[sl]]
        } else {
            sum -= m[s[sl]]
        }
        sl--
    }
    return sum
}