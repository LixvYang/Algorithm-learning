// Package leetcode provides 
package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var trace []int
	var result [][]int
	backtree(0,0,target,candidates,trace,&result)
	return result
}

func backtree(Start, sum int, target int, candidates []int, trace []int, result *[][]int) {
	if sum == target {
			tmp := make([]int, len(trace))
			copy(tmp, trace)
			*result = append(*result, trace)
			return
	}
	if sum > target {
			return
	}

	for i := Start;i<len(candidates);i++ {
			trace = append(trace, candidates[i])
			sum+=candidates[i]
			backtree(i,sum,target,candidates,trace,result)
			sum-=candidates[i]
			trace=trace[:len(trace)-1]
	}
}