package withgo
func trap(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }

    ans := 0
    p := make([]int, n)

    
    for i := 0; i < n; i++ {
        if i > 0 {
            p[i] = p[i-1]
        }
        if height[i] > p[i] {
            p[i] = height[i]
        }
    }

    
    suf := 0
    for i := n - 1; i >= 0; i-- {
        if height[i] > suf {
            suf = height[i]
        }
        min := p[i]
        if suf < min {
            min = suf
        }
        ans += min - height[i]
    }

    return ans
}