// 机器人的运动范围
package offer

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, visited)
}

func dfs(i, j, m, n, k int, visited [][]bool) int {
	if i<0 || i>=m || j<0 || j>= n || visited[i][j] || (i/10+i%10+j/10+j%10 > k) {
		return 0
	}
	visited[i][j] = true
	return dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited) + 1
}