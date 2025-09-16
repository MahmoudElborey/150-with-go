package withgo

func majorityElement(nums []int) int {
    sort.Ints(nums);
    var arrayLength int = len(nums)
    return nums[arrayLength /2]
    
}