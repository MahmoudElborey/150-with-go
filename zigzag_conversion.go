package withgo

import (
	"fmt"
	"sort"
)
func makeRange(n int) []int {
    res := make([]int, n)
    for i := range res {
        res[i] = i
    }
    return res
}

func solve(i , n , ladder , rows int, up bool , streak int , mp map[int]int) {
    if i >= n {
        return
    }

    mp[i] = ladder

   fmt.Println("i is %v and ladder is %v" , i , ladder)


    if  (up && streak == rows) || (!up && streak == rows -2)  {
        if !up {
            ladder = 0
        }
        up = !up
        streak = 0
    
    }
   

    if up {
        solve(i+1 , n , ladder + 1 , rows , up , streak + 1 ,mp)
    } else {

        solve(i+1 , n , ladder -1 , rows , up , streak+1 , mp)
    }
    return

}


func convert(s string, numRows int) string {
     n:= len(s)
    mp:=make(map[int]int)
   
    arr:= makeRange(n)
    if numRows == 1 {
        return s
    }else if (numRows == 2){
        for i:= 0; i < n; i++ {
            if i % 2 == 0 {
                mp[i] = 1
            }else {
                mp[i] = 2
            }
        }

    }else {
        solve(0 , n , 1 , numRows , true , 1 , mp)
    }
   

    sort.Slice(arr ,func(i , j int) bool {
         a, b := arr[i], arr[j]
        if mp[a] == mp[b] {
            return a < b
        }
        return mp[a] < mp[b]
    })
    var bytes []rune
    shit:= []rune(s)
    for i:= 0; i < n; i++{
        idx := arr[i]
        fmt.Println(idx , mp[idx])
       bytes = append(bytes ,shit[idx])
    }
    ans:=string(bytes)
    return ans
 
}