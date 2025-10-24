package withgo

// -1 not visited

// 1 visited and true

// 0 visted and false

func solveDp(i, n int , s string , words [] string , dp []int) int {
    if i == n {
        return 1
    }

    if dp[i] != -1 {
        return dp[i]
    }
    dp[i] = 0
    for _ , word := range words {
        wordLen := len(word)
        ch:=0
        j:= i+ wordLen
        if i+wordLen <= n && word == s[i: j] {
            ch|= solveDp(j , n, s , words , dp)
            if ch > 0 {
                dp[i] = 1
                return dp[i]
            }
        }
       
    }
    return dp[i]
}


func wordBreak(s string, wordDict []string) bool {
    n:=len(s)
    dp:=make([]int , n)
    for i:= 0; i < n; i++ {
        dp[i] = -1
    }
    can:= solveDp(0 , n ,  s , wordDict , dp)
    if can > 0 {
        return true
    }
    return false
    
}

