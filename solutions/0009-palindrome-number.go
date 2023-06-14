// Day 7

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	count := 0
	dummy := x
	for dummy != 0 {
		dummy -= (dummy % int(math.Pow10(count + 1)))
		count++
	}
	totolCount := count
	for i := count; i > 0; i-- {
		left := (x % int(math.Pow10(i))) / int(math.Pow10(i - 1))
		right := (x % int(math.Pow10(totolCount - i + 1))) / int(math.Pow10(totolCount - i))
		if right != left {
			return false
		}
	}
	return true
}