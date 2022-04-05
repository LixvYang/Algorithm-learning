// 旋转图像
package leetcode

func rotate(matrix [][]int)  {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		reverse(row)
	}
}

func reverse(n []int) {
	l, r := 0, len(n)-1
	for l < r {
		n[l], n[r] = n[r], n[l]
		l++
		r--
	}
}