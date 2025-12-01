class Solution {
    public int uniquePathsWithObstacles(int[][] grid) {
        int n = grid.length;
        int m = grid[0].length;
        int [][] dp = new int[n][m];
        
        for (int i = n-1; i >= 0; --i){
            for (int j = m-1; j >= 0; --j){
                if (i == n-1 && j == m-1){
                    dp[i][j] = 1;
                }else {
                    int ways = 0;
                    if (i+1 < n && grid[i+1][j] == 0){
                        ways+= dp[i+1][j];
                    }
                    if (j+1 < m && grid[i][j+1] == 0){
                        ways+= dp[i][j+1];
                    }
                    dp[i][j] = ways;
                }
            }
        }

        return Math.max(grid[0][0], grid[n-1][m-1]) == 1 ? 0 : dp[0][0];
        
    }
}