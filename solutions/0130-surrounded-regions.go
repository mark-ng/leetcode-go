// Day 11

func solve(board [][]byte)  {
    for i := 0; i < len(board); i++ {
        if i == 0 || i == len(board) - 1 {
            for j := 0; j < len(board[0]); j++ {
                dfs(board, i, j, 'Z')
            }
        } else {
            dfs(board, i, 0, 'Z')
            dfs(board, i, len(board[0]) - 1, 'Z')
        }
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(board, i, j, 'X')
        }
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == 'Z' {
                board[i][j] = 'O'
            }
        }
    }
}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func dfs(board [][]byte, i int, j int, replacement byte) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
        return
    }
    if board[i][j] == 'X' || board[i][j] == 'Z' {
        return
    }
    board[i][j] = replacement
    for k := 0; k < 4; k++ {
        dfs(board, i + dy[k], j + dx[k], replacement)
    }
}