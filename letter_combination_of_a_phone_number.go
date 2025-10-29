package withgo

import "fmt"



var values = []string{
    "dsf",
    "fds",
    "abc",
    "def",
    "ghi",
    "jkl",
    "mno",
    "pqrs",
    "tuv",
    "wxyz",
}

var ans []string

func solve(i int , s string , ss string)  {
    //fmt.Println("the value of i  and ss" , i , ss)
    if i == len(s){
        ans = append(ans , ss)
        return
    }

    index := int(s[i]-'0')
    fmt.Println(index)

    for _ , ch := range values[index] {
        
        solve(i+1 , s , ss + string(ch))
    }
}

func letterCombinations(digits string) []string {
    ans =nil
    solve(0 , digits , "")

    return ans
    
}