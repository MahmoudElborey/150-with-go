package withgo
func twoSum(numbers []int, target int) []int {
    n:=len(numbers)

    i:=0
    j:= n-1

    for i < j {
        curSum := numbers[i] + numbers[j]
        if curSum == target {
            return [] int {i+1 , j+1}
        }else if (curSum > target) {
            j--
        }else {
            i++
        }
    }
    return nil
    
}