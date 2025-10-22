package withgo


func verfiy (mpS , mpT map[rune]int ) bool {
    for key , value := range mpT {
        if mpS[key] < value {
            return false;
        }
    }
    return true
}

func minWindow(s string, t string) string {
    // i need substring from s
    n:= len(s)

    mpS:= make(map[rune]int)
    mpT:=make(map[rune]int)

    sRunes:=[]rune(s)
    tRunes:=[]rune(t)

    for _,ch :=range tRunes {
        mpT[ch]+=1
    }

    L:=0
    ansL := -1
    ansR:= -1

    for R:= 0; R < n; R++ {
        ch:= sRunes[R]
        mpS[ch]+=1

        for L <= R && verfiy(mpS , mpT) {

            if ansL != -1 && ansR - ansL > R-L {
                ansL = L
                ansR = R
            }else if ansL == -1 {
                ansL = L
                ansR = R
            }

            ch:= sRunes[L]
            mpS[ch]-=1
            L++
        }
    }
    
    if ansL == -1 {
        return ""
    }
    return s[ansL : ansR+1]
    
}