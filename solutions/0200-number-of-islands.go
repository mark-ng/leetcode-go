// Day 6

var dx []int = []int{0, 1, 0, -1}
var dy []int = []int{1, 0, -1, 0}

func numIslands(grid [][]byte) int {
    m, n, c := len(grid), len(grid[0]), 0
    fmt.Println(m, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                c++
                dfs(grid, i, j, m, n)
            }
        }
    }
    return c
}

func dfs(grid [][]byte, i, j, m, n int) {
    if i < 0 || i >= m || j < 0 || j >= n {
        return
    }
    if grid[i][j] == '2' || grid[i][j] == '0' {
        return
    }
    if grid[i][j] == '1' {
        grid[i][j] = '2'
        for k := 0; k < 4; k++ {
            dfs(grid, i + dy[k], j + dx[k], m, n)
        }
    }
}