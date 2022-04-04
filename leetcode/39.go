// 组合总和
package leetcode

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	backtrack(0, 0, target, &res, []int{}, candidates)
	return res
}

func backtrack(startIndex, sum, target, int, res *[][]int, track, candidates []int) {
	if sum > target {
		return 
	}

	if sum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	for i := startIndex; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack(i, sum, target, res, track, candidates)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}