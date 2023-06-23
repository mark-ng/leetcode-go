// Day 12

type Stack struct {
    inside []string
}

func (stk *Stack) push(item string) {
    stk.inside = append(stk.inside, item)
}

func (stk *Stack) pop() (string, error) {
	if len(stk.inside) == 0 {
		return "", errors.New("Nothing to pop")
	}
	popV := stk.inside[len(stk.inside) - 1]
	stk.inside = stk.inside[:len(stk.inside) - 1]
	return popV, nil
}

func (stk *Stack) isEmpty() bool {
    return len(stk.inside) == 0
}

func simplifyPath(path string) string {
    stk := Stack{
        inside: []string{},
    }
    ps := strings.Split(path, "/")
    
    // Clean ps
    newps := []string{}
    for _, p := range ps {
        if p == " " || p == "" || p == "/" {
            continue
        } 
        newps = append(newps, p)
    }

    for _, p := range newps {
        if p == "." {
            continue
        } else if p == ".." {
            stk.pop()
        } else {
            stk.push(p)
        }
    }

    var ans strings.Builder
    for i:= 0; i < len(stk.inside); i++ {
        ans.WriteString("/" + stk.inside[i])
    }

    if ans.String() == "" {
        return "/"
    }

    return ans.String()
}