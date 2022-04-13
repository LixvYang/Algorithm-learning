// 单词搜索
package offer

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var canFind func(r, c, i int) bool
	canFind = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r<0 || r>=m || c<0 || r>=n || used[r][c] {
			return false
		}

		if board[r][c] != word[i] {
			return false
		}

		used[r][c] = true
		if canFind(r-1, c, i+1) || canFind(r+1, c, i+1) || canFind(r, c-1, i+1) || canFind(r, c+1, i+1) {
			return true
		} else {
			used[r][c] = false
			return false
		}	
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && canFind(i, j, 0) {
				return true
			}
		}
	}
	return false
}