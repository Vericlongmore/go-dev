package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i != -1; i-- {
		digits[i]++
		if digits[i]/10 == 0 {
			return digits
		}
		digits[i] = digits[i] % 10
	}
	return append([]int{1}, digits...)
}
