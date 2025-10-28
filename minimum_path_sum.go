package withgo



func minInt(a , b int) int {
    if a <= b {
        return a
    }

    return b
}

func minPathSum(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])
    dp:= make([][]int , n)
    for i:= 0; i < n; i++ {
        dp[i] = make([] int , m)
    }
    for i:= 0; i < n; i++ {
        for j:= 0; j < m; j++ {
            dp[i][j] = 200 * 200 * 200
        }
    }

    for i:= n-1; i >= 0; i-- {
        for j:=m-1; j>= 0; j-- {
            if i == n-1 && j == m-1 {
                dp[i][j] = grid[i][j];
            }
            if i+1 < n {
                dp[i][j] = minInt(dp[i][j] , grid[i][j] + dp[i+1][j])
            }

            if j+1 < m {
                dp[i][j] = minInt(dp[i][j] , grid[i][j] + dp[i][j+1])
            }
        }
    }
    return dp[0][0]
}