// Day 8

func reverseBits(num uint32) uint32 {
    ans := uint32(0)
    count := 32
    for count != 0 {
        ans = ans << 1
        b := num & 0x1
        ans = ans + b
        num = num >> 1
        count--
    }
    return ans
}