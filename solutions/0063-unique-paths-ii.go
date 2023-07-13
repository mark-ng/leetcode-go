// Day 22

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

    m, n := len(obstacleGrid), len(obstacleGrid[0])

    ans := make([][]int, m)
    for i := 0; i < m; i++ {
        ans[i] = make([]int, n)
    }
    
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    ans[0][0] = 1

    for j := 1; j < n; j++ {
        if obstacleGrid[0][j] == 1 {
            ans[0][j] = 0
        } else {
            ans[0][j] = ans[0][j - 1]
        }
    }
    for i := 1; i < m; i++ {
        if obstacleGrid[i][0] == 1 {
            ans[i][0] = 0
        } else {
            ans[i][0] = ans[i - 1][0]
        }
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                ans[i][j] = 0
                continue
            }
            up := ans[i - 1][j]
            left := ans[i][j - 1]
            ans[i][j] = up + left
        }
    }
    return ans[m - 1][n - 1]   
}