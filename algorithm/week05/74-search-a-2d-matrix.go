func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1 // 转一维
	for l <= r {
		mid := (r-l)>>1 + l // mid是一维的索引

		curRow := mid / cols // 整除，得二维的当前行索引
		curCol := mid - curRow*cols // 一维mid减去它头顶上行的元素个数，得二维的当前列索引

		if matrix[curRow][curCol] == target {
			return true
		} else if matrix[curRow][curCol] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false

}
