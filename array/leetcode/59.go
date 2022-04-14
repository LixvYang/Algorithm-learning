// 螺旋矩阵2
package leetcode

func generateMatrix(n int) [][]int {
	martix := make([][]int, n)
	for i := range martix {
		martix[i] = make([]int, n)
	}

	top, left, right, bottom := 0, 0, n-1, n-1
	// 	
	num, target := 1, n*n

	for num <= target {
		for i:=left; i <= right; i++ {
			martix[top][i] = num
			num++
		}
		top++
		for i:=top;i<=bottom;i++ {
			martix[i][right] = num
			num++
		}
		right--
		for i:=right;i>=left;i-- {
			martix[bottom][i] = num
			num++
		}
		bottom--
		for i:=bottom;i>=top;i-- {
			martix[i][left] = num
			num++
		}
		left++
	}
	return martix	
}