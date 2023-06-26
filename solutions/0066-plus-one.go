// Day 15

func plusOne(digits []int) []int {
    digitsLen := len(digits)
    newDigits := make([]int, digitsLen + 1) 
    
    carry := 0 
    for po, pn := digitsLen - 1, digitsLen; po >= 0; po, pn = po - 1, pn - 1 {
        var v int
        if po == digitsLen - 1 {
            v = digits[po] + 1 + carry
        } else {
            v = digits[po] + carry
        }
        if v >= 10 {
            carry = 1
        } else {
            carry = 0
        }
        newDigits[pn] = v % 10
    }
    if carry == 1 {
        newDigits[0] = 1
        return newDigits[:]
    }
    return newDigits[1:]

}