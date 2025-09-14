package withgo
import (
    "fmt"
)

func shiftElementsRight(nums1 [] int , n int , finalArrayLength int){
    for i:= finalArrayLength - 1 ; i >= n ; i-- {
        nums1[i] = nums1[i-n]
    }
}
func merge(nums1 []int, m int, nums2 []int, n int)  {
    finalArrayLength := n + m
    
    mergedArrayIndex , secondArrayIndex :=n, 0

    shiftElementsRight(nums1 , n , finalArrayLength)
    

    for i:=0; i < finalArrayLength ; i++ {
        if mergedArrayIndex >= 0 && mergedArrayIndex < finalArrayLength && secondArrayIndex < n {
            if  nums1[mergedArrayIndex] <= nums2[secondArrayIndex] {
               nums1[i] = nums1[mergedArrayIndex]
               mergedArrayIndex++
            } else {
                nums1[i] = nums2[secondArrayIndex]
                secondArrayIndex++;
            }
        } else if mergedArrayIndex < 0 || mergedArrayIndex >= finalArrayLength {
            nums1[i] = nums2[secondArrayIndex]
            secondArrayIndex++;
        }else {
            nums1[i] = nums1[mergedArrayIndex]
            mergedArrayIndex++
        }
       
    }
}