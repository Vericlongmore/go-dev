func shipWithinDays(weights []int, D int) int {
	//1 <= D <= weights.length <= 50000
	//1 <= weights[i] <= 500
	//left，right是 船载重范围
	left, right := 1, 500*len(weights)/D
	//二分缩小载重范围
	for left <= right {
		mid := left + (right-left)/2
		if shipWithinDaysValidate(weights, D, mid) {
			//满足缩小上边界
			right = mid - 1
		} else {
			//不满足增大下边界
			left = mid + 1
		}
	}
	return left
}
//验证size大小的载重是否满足情况，满足返回true
func shipWithinDaysValidate(weights []int, D, size int) bool {
	count := 1
	has := 0
	for i := range weights {
		if weights[i] > size {
			return false
		}
		if has+weights[i] <= size {
			has += weights[i]
		} else {
			has = weights[i]
			count++
		}
	}
	return count <= D
}
