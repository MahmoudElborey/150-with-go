package withgo


func rotate(nums []int, k int)  {
    var length = len(nums)
    k%=length
    tempArray := make( []int, length)
    for i:= 0 ; i < length ; i++ {
        var nextIndex int = (i+ k) %length
        tempArray[nextIndex] = nums[i]
    }
    for i:=0; i < length; i++ {
        nums[i] = tempArray[i]
    }
    
}