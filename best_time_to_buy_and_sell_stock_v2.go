package withgo

import "fmt"

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}



func maxProfit(prices []int) int {
    var length int = len(prices)
    maxAnswer , maxDp := 0 , 0
    dp:= make([]int , length+1)
    for i:= length -1; i >= 0 ; i-- {
        dp[i] = - prices[i] + maxDp
        dp[i] = maxInt(dp[i+1] , dp[i])
        maxDp = maxInt(maxDp , dp[i+1] + prices[i])
        maxAnswer = maxInt(maxAnswer , dp[i])
        //fmt.Println("index i %d and it's value %d and maxDp is %d" , i , dp[i] , maxDp)
    }
    return maxAnswer
}