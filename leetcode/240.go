//搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	// 二分查找，搜索二维矩阵中的值，左至右递增，上到下递增

	rows, cols := len(matrix), len(matrix[0])
	// 初始化到最左下角的元素位置
	// 因为该位置元素大小处于中间值，相对来说锯齿线性更合理
	i, j := rows-1, 0
	// i只能减，j只能增，进而转化为二分查找，i相当于右边界，j是左边界
	// matrix[i][j]->mid
	for i >=0 && j < cols {
		// target大于该值，需要增大mid，或者更新左边界
		if matrix[i][j] < target {
			j++
		// target小于该值，需要减小mid，或者更新右边界
		} else if matrix[i][j] > target {
			i--
		// 找到该元素，直接返回true
		} else {
			return true
		}
	}
	// 锯齿形遍历没有找到，false
	return false
}