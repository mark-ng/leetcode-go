// Day 10

func canConstruct(ransomNote string, magazine string) bool {
    h := map[byte]int{}
    for i := 0; i < len(magazine); i++ {
        h[magazine[i]] += 1
    }
    for i := 0; i < len(ransomNote); i++ {
        v := h[ransomNote[i]]
        if v == 0 {
            return false
        }
        h[ransomNote[i]]--
    }
    return true
}