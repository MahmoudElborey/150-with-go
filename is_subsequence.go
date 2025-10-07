package withgo

func solve(i , j , sLength , tLength int , s string , t string) bool {
    if (i == sLength) {
        return true
    }
    if (j == tLength) {
        return false
    }

    if (s[i] == t[j]) {
        return solve(i+1 , j+1 , sLength , tLength , s , t)
    }
    return solve(i , j+1 , sLength , tLength ,s , t)
}

func isSubsequence(s string, t string) bool {
    return solve(0 , 0 , len(s) , len(t) , s , t)
    
}