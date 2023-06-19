// Day 9

func hammingWeight(num uint32) int {
    count := 0
    for num != 0 {
        b := num & 0x1
        if b == 1 {
            count++
        }
        num = num >> 1
    }
    return count
}