func removeDuplicates(nums []int) int {
    arrayLength:= len(nums)
    var answer int = 1
    sort.Ints(nums)
    

    for i:= 1; i < arrayLength; i++ {
        if nums[i] != nums[i-1] {
            nums[answer] = nums[i]
            answer++
        }
    }
    return answer
}