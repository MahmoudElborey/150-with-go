package withgo
func maxInt(a int , b  int) int {
    if a < b {
        return b
    }
    return a
}


func hIndex(citations []int) int {
    arraySize := len(citations)
    freq := make([] int , 5002)
    for i:=0; i < arraySize ; i++ {
        freq[citations[i]]++
    }
    maxCandidate := 0
    for i:= 5000; i >= 0 ; i-- {
        freq[i]+= freq[i+1]
        if freq[i] >= i {
            maxCandidate = maxInt(maxCandidate , i)
        }
    }
    return maxCandidate
    
}