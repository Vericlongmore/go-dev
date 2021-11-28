func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
func findShortestSubArray(nums []int) int {
	res := 0
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
		res = max(res, hash[v])
	}
	n := len(nums)
	has := make(map[int]int)
    hash2 := make(map[int]int)
	ans := n
    if res == 1{
        return 1
    }
	for i := 0; i < n; i++ {
        hash2[nums[i]]++
		if idx, ok := has[nums[i]]; ok {
			if hash2[nums[i]] == res {
                ans = min(ans, i - idx + 1)
            }
		}else{
            has[nums[i]] = i
        }
	}
	return ans
}
