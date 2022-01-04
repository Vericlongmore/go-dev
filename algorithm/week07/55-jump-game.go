func canJump(nums []int) bool {
    if len(nums) == 0{
        return false
    }
    canReachable := len(nums) -1
    for i:= canReachable; i >=0 ;i--{
        if nums[i] + i >= canReachable{
            canReachable = i
        }
    }

    return canReachable == 0


}
