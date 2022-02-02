// Package leetcode provides 
package leetcode

func combinationSum2(candidates []int, target int) [][]int {
	var track []int
	var res [][]int
	var history map[int]bool
	history=make(map[int]bool)
	sort.Ints(candidates)
	backtracking(0,0,target,candidates,track,&res,history)
	return res
}

func backtracking(startIndex, sum int,target int, candidates []int, track []int, res *[][]int, history map[int]bool) {
	if sum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			*res = append(*res, tmp)
			return
	}

	if sum > target{
			return
	}

	for i := startIndex; i < len(candidates);i++ {
			if i>0 && candidates[i]==candidates[i-1]&&history[i-1]==false {
					continue
			}
			track = append(track, candidates[i])
			sum+=candidates[i]
			history[i]=true
			backtracking(i+1,sum,target,candidates,track,res,history)
			track=track[:len(track)-1]
			sum-=candidates[i]
			history[i]=false
	}
}