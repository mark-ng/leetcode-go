// Day 13

var phoneMap = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

func letterCombinations(digits string) []string {
    ans := []string{}
    if len(digits) == 0 {
        return ans
    }
    dfs(digits, 0, &ans, "")
    return ans
}

func dfs(digits string, idx int, ans *[]string, cur string) {
    if idx == len(digits) {
        *ans = append(*ans, cur)
        return
    }
    n := string(digits[idx])
    for i := 0; i < len(phoneMap[n]); i++ {
        cur += string(phoneMap[n][i])
        dfs(digits, idx + 1, ans, cur)
        cur = cur[:len(cur) - 1]
    }
}