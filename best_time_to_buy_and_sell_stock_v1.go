package withgo


func maxInt(a int , b int ) int {
    if (a < b){
        return b
    }
    return a
}

func minInt (a int , b int) int {
    if (a > b){
        return b
    }
    return a
}

func maxProfit(prices []int) int {
    var length int = len(prices)
    prefixMinimum:=prices[0]
    var maxAnswer int = 0
    for i:=1; i < length; i++ {
        maxAnswer = maxInt(maxAnswer , prices[i] - prefixMinimum)
        prefixMinimum = minInt(prefixMinimum , prices[i])
    }
    return maxAnswer
    
}