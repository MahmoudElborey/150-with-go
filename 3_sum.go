package withgo

import (
	"sort"
)


func threeSum(nums []int) [][]int {

    sort.Ints(nums)

    var ans [][]int

    n:= len(nums)

    for i:= 0; i+ 2< n; i++ {
        R:= n-1
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j:=i+1; j +1  < n; j++ {

            if nums[j] == nums[j-1] && i +1 != j {
                continue
            }

            oth:= -1* (nums[i] + nums[j])

            for R > j && nums[R] > oth {
                R--
            }

            if R > j && oth == nums[R] {
                curAns:= []int {nums[i] , nums[j] , oth}
                ans = append(ans , curAns)
            }



            
            
        }
    }

    return ans
    
}