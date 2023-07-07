// Day 19

func evalRPN(tokens []string) int {
    var stk []string 
    for i := 0; i < len(tokens); i++ {
        token := tokens[i]
        switch token {
            case "+":
                a := stk[len(stk) - 1]
                stk = stk[:len(stk) - 1]
                b := stk[len(stk) - 1]   
                stk = stk[:len(stk) - 1]
                aa, _ := strconv.Atoi(a) 
                bb, _ := strconv.Atoi(b) 
                result := strconv.Itoa(bb + aa)
                stk = append(stk, result)
            case "-":
                a := stk[len(stk) - 1]
                stk = stk[:len(stk) - 1]
                b := stk[len(stk) - 1]   
                stk = stk[:len(stk) - 1]
                aa, _ := strconv.Atoi(a) 
                bb, _ := strconv.Atoi(b) 
                result := strconv.Itoa(bb - aa)
                stk = append(stk, result)
            case "*":
                a := stk[len(stk) - 1]
                stk = stk[:len(stk) - 1]
                b := stk[len(stk) - 1]   
                stk = stk[:len(stk) - 1]
                aa, _ := strconv.Atoi(a) 
                bb, _ := strconv.Atoi(b) 
                result := strconv.Itoa(bb * aa)
                stk = append(stk, result)
            case "/":
                a := stk[len(stk) - 1]
                stk = stk[:len(stk) - 1]
                b := stk[len(stk) - 1]   
                stk = stk[:len(stk) - 1]
                aa, _ := strconv.Atoi(a) 
                bb, _ := strconv.Atoi(b) 
                result := strconv.Itoa(bb / aa)
                stk = append(stk, result)
            default:
                stk = append(stk, token)
        }
    }
    ans, _ := strconv.Atoi(stk[0]) 
    return ans 
}