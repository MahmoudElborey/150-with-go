package withgo
func lengthOfLastWord(s string) int {
    lengthOfString:= len(s)
    ans:= 0
    start:=false

    for i:= lengthOfString -1; i >= 0; i-- {
        if s[i] == ' ' {
            if start {
                break
            }
            continue
        }
        ans++
        start = true
    }
    return ans
    
}