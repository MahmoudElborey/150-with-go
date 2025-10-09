package withgo

func lengthOfLongestSubstring(s string) int {
    last:= make(map[rune]int)

    ans:=0
    lastIndex:=0
    runes:= []rune(s)

    for i , ch := range runes {

        if idx , found := last[ch] ; found {
            
            if (idx + 1 > lastIndex){
                lastIndex = idx + 1
            }
        }

        if i - lastIndex +1 > ans {
            ans = i - lastIndex + 1
        }

        last[ch] = i
    }

    return ans
    
}

