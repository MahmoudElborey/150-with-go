package withgo

import "fmt"
func maxInt (a , b int) int {
    if (a < b){
        return b
    }
    return a
}


func solveCase1(h []int , n int) int{
    sufMax:= make([]int , n)
    sufMax[n-1] = h[n-1]
    for i:= n-2; i >=0 ; i-- {
        sufMax[i] = sufMax[i+1]
        sufMax[i] = maxInt(sufMax[i] , h[i])
    }

    result:= 0

    for i:= 0; i < n; i++ {
        curAns := -1
        L:= i+1
        R:= n-1

        for L <= R {
            md := int(L + R) / 2
            if h[i] > sufMax[md] {
               R = md -1
            }else {
                curAns = md
                L = md + 1
            }
        }

        fmt.Println("i is %v and curAns is %v", i , curAns)

        if curAns != -1 {
            result = maxInt(result , (curAns -i) * h[i])
        }

    }

    return result
}

func solveCase2(h [] int , n int) int {
    result := 0
    preMax := make([] int , n)
    preMax[0] = h[0]
    for i:=1 ;i < n; i++ {
        preMax[i] = preMax[i-1]
        preMax[i] = maxInt(preMax[i] , h[i])
    }

    for i:=0; i < n; i++ {
        curAns := -1
        L := 0
        R:= i-1

        for L <= R {
            md := (L+ R) / 2
            if h[i] >= preMax[md] {
                L = md + 1
            }else {
                R = md - 1
                curAns = md
            }
        }
        if curAns != -1 {
            result = maxInt(result , (i - curAns) * h[i])
        }
    }
    return result
}

func maxArea(height []int) int {
    n:= len(height)
    result := solveCase1(height , n)

    result = maxInt(result ,solveCase2(height , n) )
    return result
}