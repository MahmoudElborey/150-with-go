package withgo

import (
	"fmt"
	"unicode"
)
func isPalindrome(s string) bool {
    var result [] rune
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            result = append(result , unicode.ToLower(r))
        }
    }
    

    n:= len(result)
    fmt.Println(n)
    for i:=0; i < n/ 2; i++ {
        fmt.Println( string(result[i]) + " shit  " + string(result[n-1-i]))
        if result[i] != result[n-i-1] {
            return false
        }
    }
    return true

    
}