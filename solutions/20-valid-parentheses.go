func isValid(s string) bool {
    stk := []rune{}
	pair := map[rune]rune{
		')':'(',
		']':'[',
		'}':'{',
	}
    for _, v := range s {
		switch v {
		case '(', '[', '{':
            stk = append(stk, v)
		case ')', ']', '}':
			if len(stk) == 0 {
				return false
			}
			popV := stk[len(stk) - 1]
			if popV != pair[v] {
				return false
			}
			stk = stk[:len(stk) - 1]
		}
    }
    return len(stk) == 0
}