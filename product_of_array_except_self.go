package withgo
func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)

    mul := 1
    for i := 0; i < n; i++ {
        ans[i] = mul
        mul *= nums[i]
    }

    mul = 1
    for i := n - 1; i >= 0; i-- {
        ans[i] *= mul
        mul *= nums[i]
    }

    return ans
}
