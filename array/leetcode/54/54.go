package leetcode

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
}

	res := make([]int, 0)
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for left <= right && top <= bottom {
			for i := left; i <= right; i++ {
					res = append(res, matrix[top][i])
			}
			top++
			for i := top; i <= bottom; i++ {
					res = append(res, matrix[i][right])
			}
			right--
			if left>right || top>bottom {
					break
			}
			for i := right; i >= left; i-- {
					res = append(res, matrix[bottom][i])
			}
			bottom--
			for i := bottom; i >= top; i-- {
					res = append(res, matrix[i][left])
			} 
			left++
	}
	return res
}