package withgo
func canCompleteCircuit(gas []int, cost []int) int {

    n:= len(gas)
    preSum := 0
    ans := 0
    tot:=0
    for i:= 0; i < n; i++ {
        value:= gas[i] - cost[i]
        preSum+= value
        tot+= value
        if preSum < 0 {
            preSum = 0
            ans = i+1
        }
    }
    if tot < 0 {
        ans = -1;
    }
    return ans
    
}