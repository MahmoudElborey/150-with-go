package withgo



func solveDp(weight int, dp , coins[]  int) int {
    if weight == 0 {
        return 0
    }

    ret := &dp[weight]

    if *ret != -1 {
        return *ret
    }

    *ret = int(1e5)


    for _ , coin := range coins {
        if coin <= weight {
            ch := 1 + solveDp(weight - coin , dp , coins)
            if *ret > ch {
                *ret = ch
            }
        }
    }
    return dp[weight]
}



func coinChange(coins []int, amount int) int {

    dp:= make([]int , amount+5)
    for i:= 0; i <= amount; i++ {
        dp[i] = -1
    }

    value := solveDp(amount , dp , coins)
    if value >= int(1e5){
        value = -1
    }
    return value
    
} 