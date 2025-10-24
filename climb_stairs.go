package withgo

func climbStairs(n int) int {

    dp:= make([]int , 50)

    dp[0] = 1

    for i:= 1; i <= 45; i++ {
        dp[i] = dp[i-1]
        if i > 1 {
            dp[i]+= dp[i-2]
        }
    }

    return dp[n]
}