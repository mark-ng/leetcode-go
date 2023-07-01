// Day 18

func searchMatrix(matrix [][]int, target int) bool {
    ask1 := func(idx int) bool {
        return matrix[idx][0] >= target
    }
    targetRow := findFirst(ask1, len(matrix))
    if targetRow > len(matrix) {
        return false
    }
    targetRow -= 1

    if targetRow + 1 < len(matrix) && matrix[targetRow + 1][0] == target {
        return true
    }
    if targetRow < 0 {
        return false
    }
    ask2 := func(idx int) bool {
        return matrix[targetRow][idx] >= target
    }
    targetCol := findFirst(ask2, len(matrix[targetRow]))
    fmt.Println(targetCol)
    if targetCol >= len(matrix[targetRow]) {
        return false
    }
    return matrix[targetRow][targetCol] == target
}

func findFirst(ask func(int) bool, n int) int {
    lo := 0
    hi := n - 1

    if n == 0 || ask(lo) == true {
        return lo 
    }
    if ask(hi) == false {
        return n
    }

    for hi - lo > 1 {
        mid := (hi + lo) / 2
        if ask(mid) == true {
            hi = mid
        } else {
            lo = mid
        }
    } 
    return hi
}

// REDO -> compressed 2d to 1d