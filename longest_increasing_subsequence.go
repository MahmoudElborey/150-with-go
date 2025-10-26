package withgo



func solveDp (i  , n int , dp ,  nums [] int ) int {
    if i == n {
        return 0
    }

    ret:= & dp[i]
    if *ret != -1 {
        return *ret
    }

    *ret = 0

    for k := i+1 ; k < n; k++ {
        if nums[k] > nums[i] {
            ch:= 1 + solveDp(k , n , dp , nums)
            if ch > * ret {
                *ret = ch
            }
        }
    }

    return *ret
}


func lengthOfLIS(nums []int) int {
    n:= len(nums)
    dp:= make([]int , n)
    for i:=0; i < n; i++ {
        dp[i] = -1
    }
    
    ans := 0
    for last := 0; last < n; last++ {
        curAns := 1 + solveDp (last , n , dp , nums)
        if curAns > ans {
            ans = curAns
        }
    }

    return ans
}