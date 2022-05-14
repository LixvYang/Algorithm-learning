// 所有可能的路径

package leetcode

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := make([][]int, 0)
n := len(graph)
var dfs func([]int, int)
dfs = func(ints []int, index int) {
	if index == n-1 {
		temp := make([]int, len(ints))
		copy(temp, ints)
		ans = append(ans, temp)
		return
	}
	for _, v := range graph[index] {
		ints = append(ints, v)
		dfs(ints, v)
		ints = ints[:len(ints)-1]
	}
}
dfs([]int{0}, 0)
return ans
}