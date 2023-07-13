// Day 22

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i - 1 < 0 && j - 1 < 0 {
                continue
            }
            if i - 1 < 0 {
                grid[i][j] = grid[i][j] + grid[i][j - 1]
                continue
            }
            if j - 1 < 0 {
                grid[i][j] = grid[i][j] + grid[i - 1][j]
                continue
            } 
            up := grid[i - 1][j]
            left := grid[i][j - 1]
            
            if up < left {
                grid[i][j] = up + grid[i][j]
            } else {
                grid[i][j] = left + grid[i][j]
            }
        }
        
    }
    return grid[m - 1][n - 1]
}