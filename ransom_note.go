package withgo


func canConstruct(ransomNote string, magazine string) bool {
    magMap:=make(map[rune]int)


    for _ , ch:= range magazine {
        magMap[ch]+=1
    }


    for _ , ch:= range ransomNote {
        magMap[ch]--

        if magMap[ch] < 0 {
            return false
        }
    }

    return true
    
}