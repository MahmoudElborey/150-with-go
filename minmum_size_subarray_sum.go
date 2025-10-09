package withgo

func minInt(a , b int) int {
    if a <= b {
        return a
    }
    return b
}

func minSubArrayLen(target int, nums []int) int {
    n:= len(nums)
    ans:= int(1e9)
    L:= 0
    curSum :=0
    for R:=0; R < n; R++ {
        curSum+= nums[R]

        for L <= R && curSum >= target {
            curSum-= nums[L]
            ans = minInt(ans , R - L +1)
            L++
        }
    }

    if ans == int(1e9) {
        ans = 0
    }
    return ans
    
}