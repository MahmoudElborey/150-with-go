package withgo
func check(stI , stJ , endI , endJ int , board[][] byte) bool {
    mp := make(map[byte]bool)
    for i:= stI; i <= endI; i++ {
        for j:= stJ ; j <= endJ; j++ {
            cur:= board[i][j]

             if cur == '.' {
                continue
            }

           if mp[cur] {
                return false
            }
            mp[cur] = true
        }
    }
    return true
}


func checkRaws(board [][] byte) bool {
    

    for i:= 0; i < 9; i++ {
        mp := make(map[byte]bool)
        for j:=0 ; j < 9; j++{
            cur := board[i][j]

            if cur == '.' {
                continue
            }

            if cur < '1' || cur > '9' { 
               
                return false
            }

            if mp[cur] {
            
                return false
            }
            mp[cur] = true
        }
    }
    return true
}

func checkCols(board [][] byte) bool {
    

    for i:= 0; i < 9; i++ {
        mp := make(map[byte]bool)
        for j:=0 ; j < 9; j++{
            cur := board[j][i]

             if cur == '.' {
                continue
            }

            if mp[cur] {
                return false
            }

            mp[cur] = true
        }
    }
    return true
}


func isValidSudoku(board [][]byte) bool {

    if !checkRaws(board) || !checkCols(board) {
        return false
    }

    for i:= 0; i < 9; i+=3 {
        for j:=0; j < 9; j+=3 {
            

            if !check(i , j , i + 2 , j + 2, board) {
                return false
            }
        }
    }



    return true
    
}