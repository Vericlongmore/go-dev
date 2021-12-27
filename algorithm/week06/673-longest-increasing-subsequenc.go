package main

func main(){

	nums:= []int{1,3,5,4,7}

	findNumberOfLIS(nums)
}


func findNumberOfLIS(nums []int) int {
	size := len(nums)
	if size <= 1  {
		return size
	}

	dp := make([]int, size)
	for i, _ := range dp {
		dp[i] = 1
	}
	count := make([]int, size)
	for i, _ := range count {
		count[i] = 1
	}

	maxCount := 0
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j] + 1 == dp[i] {
					count[i] += count[j]
				}
			}
			if dp[i] > maxCount {
				maxCount = dp[i]
			}
		}
	}

	result := 0
	for i := 0; i < size; i++ {
		if maxCount == dp[i] {
			result += count[i]
		}
	}
	return result
}
