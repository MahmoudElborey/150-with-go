package withgo

func removeDuplicates(nums []int) int {
    arrayLength := len(nums)
    var frequencyOfCurrentElement int = 0
    index := 1

    for i:=1 ; i < arrayLength ; i++ {
        if nums[i] == nums[i-1] {
            frequencyOfCurrentElement++
        }else {
            if (frequencyOfCurrentElement >= 1){
                nums[index] = nums[i-1]
                index++
            }

            frequencyOfCurrentElement = 0
            nums[index] = nums[i]
            index++
        }
    }
    if (frequencyOfCurrentElement >= 1){
        nums[index] = nums[arrayLength-1]
        index++
    }
    return index
}