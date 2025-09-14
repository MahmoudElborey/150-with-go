package withgo
func removeElement(nums []int, val int) int {
    arrayLength := len(nums)
    var pos int = 0
    for i:= 0; i < arrayLength; i++ {
        if nums[i] != val {
            nums[pos] = nums[i]
            pos++
        }
        
    }
    
    return pos

}