// Day 4

func addBinary(a string, b string) string {

	al, bl := len(a), len(b)
	maxl := max(al, bl) - 1
	if al > bl {
		for i := 0; i < al-bl; i++ {
			b = "0" + b
		}
	} else {
		for i := 0; i < bl-al; i++ {
			a = "0" + a
		}
	}

	carry := 0
	ans := ""
	for maxl >= 0 {
		aInt, err := strconv.Atoi(string(a[maxl]))
		if err != nil {
			return ""
		}
		bInt, err := strconv.Atoi(string(b[maxl]))
		if err != nil {
			return ""
		}
		v := (aInt + bInt + carry) % 2
		ans = strconv.Itoa(v) + ans
		carry = (aInt + bInt + carry) / 2
		maxl--
	}

	if carry == 1 {
		return "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
