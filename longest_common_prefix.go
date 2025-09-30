package withgo

import "strings"


func longestCommonPrefix(strs []string) string {
    //arrayLength := len(strs)
    lengthOfFirstString := len(strs[0])
  
    for i:= lengthOfFirstString; i > 0 ; i-- {
        can := true

        for _ , word := range strs {
            if len(word) < i || !strings.HasPrefix(word , strs[0][:i]){
                can = false 
                break
            }
        }

        if can {
            return strs[0][:i]
        }

    }
    return ""
    
}