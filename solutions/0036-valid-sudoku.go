// Day 21

func isValidSudoku(board [][]byte) bool {
    for _, row := range board {
        if !isValidRow(row) {
            return false
        }
    }
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if i % 3 == 0 && j % 3 == 0 {
                if !isValidBox(board, i, j) {
                    return false
                }
            }
        }
    }
    for i := 0; i < 9; i++ {
        temp := []byte{}
        for j := 0; j < 9; j++ {
            temp = append(temp, board[j][i])
            if !isValidRow(temp) {
                return false
            }
        }
    }
    return true
}

func isValidRow(row []byte) bool {
    set := [10]bool{}
    for _, v := range row {
        if v == '.' {
            continue
        }
        pos := v - '0'
        if set[pos] == true {
            return false
        } else {
            set[pos] = true
        }
    }
    return true
}

func isValidBox(board [][]byte, row int, col int) bool{
    set := [10]bool{}
    for i := row; i < row + 3; i++ {
        for j := col; j < col + 3; j++ {
            if board[i][j] == '.' {
                continue
            }
            pos := board[i][j] - '0'
            if set[pos] == true {
                return false
            } else {
                set[pos] = true
            }
        }
    }
    return true
}