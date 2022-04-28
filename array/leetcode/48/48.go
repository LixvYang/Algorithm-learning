// 旋转图像
package leetcode


func Rotate(matrix [][]int) [][]int {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}	

	for _, row := range matrix {
		for i := 0; i < n/2; i++ {
			row[i], row[n-i-1] = row[n-i-1], row[i]
		}
	}
	return matrix
}