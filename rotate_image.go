package withgo

func rotate(matrix [][]int)  {
    //j , n - i -1
    n:= len(matrix)
    ans:=make([][]int , n)
    for i:= 0; i < n; i++ {
        ans[i] = make([]int , n)
    }

    for i:= 0; i < n; i++ {
        for j:= 0; j < n; j++ {
            ans[j][n-i-1] = matrix[i][j]
        }
    }

    for i:= 0; i < n; i++ {
        for j:= 0;j < n; j++ {
            matrix[i][j] = ans[i][j]
        }
    }

    
}