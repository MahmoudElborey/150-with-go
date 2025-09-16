package withgo

func maxInt(a int , b int) int {
    if a > b {
        return a
    }
    return b
}

func canJump(nums []int) bool {
    var arrayLength int = len(nums)
    curMax:= nums[0]
    for i:=1; i < arrayLength; i++ {
        if curMax >= i {
            curMax = maxInt(curMax , i + nums[i])
        }
    }
    return curMax >= arrayLength -1
    
}