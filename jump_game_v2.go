package withgo
func maxInt(a int , b  int) int {
    if a < b {
        return b
    }
    return a
}

func jump(nums []int) int {
    n := len(nums)

    var jumps int = 0;
    var farthest int = 0;
    var untill int = 0;

    for i:= 0 ; i < n -1; i++ {

        farthest = maxInt(farthest , i + nums[i])

        if (i == untill){
            jumps++
            untill = farthest
        }
    }
    return jumps
    
}