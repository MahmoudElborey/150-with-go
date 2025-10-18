package withgo

import (
	"fmt"
	"os"
)

// right , down , left , up -> untill i visit the starting cell
//  0       1       2     3
// when i find the starting cell
// go down by one step and right by one step and start collecting the elements again

func solve(i , j , dir , stR , stC  , n , m  , idx int , matrix [][] int , ansList [] int , vis[][]bool) {
    if idx == n * m {
        return 
    }

    fmt.Println("value of i %v and value of j %v and stR %v andd stC %v and idx %v " , i , j , stR , stC , idx)

    if i == stR && j == stC && vis[i][j] {
        solve(i+1 , j+1 , 0, i+1 , j+1 , n , m , idx , matrix , ansList , vis)
        return
    }
   
    ansList[idx] = matrix[i][j]
    vis[i][j] = true
   
    idx++

    if dir == 0 { // right then down
        j++
        if j >= m || vis[i][j]{
            j--
            solve(i+1 , j , 1 , stR , stC , n , m , idx , matrix , ansList , vis)
            return
        }
    }else if dir == 1{ // down then left
        i++
        if i >= n || vis[i][j]{
            i--
            solve(i , j-1 , 2 , stR , stC , n , m , idx , matrix , ansList , vis)
            return
        }
    }else if dir == 2 { // left then up
        j--
        if j < 0 || vis[i][j]{
            j++
            solve(i-1 ,  j, 3 , stR , stC , n , m , idx , matrix , ansList , vis)
            return
        }
    }else {
        i--
        if i < 0 {
            os.Exit(1)
        }
    }

    solve(i , j , dir , stR , stC , n , m , idx , matrix , ansList , vis)
    return




}


func spiralOrder(matrix [][]int) []int {
    n:= len(matrix)
    m:= len(matrix[0])
    fmt.Println("value of n %v and value of m %v " , n , m)

    ansList:= make([]int , n * m)

    vis := make([][]bool, n)      
    for i := range vis {
        vis[i] = make([]bool, m)  
    }

    solve(0 , 0 , 0 , 0 , 0 ,  n , m , 0, matrix , ansList , vis)

    return ansList
    
}