// Package leetcode provides 
package leetcode
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	track := []int{}
	quickSort(candidates)
	backtrack(0, 0, target, track, candidates, &res)
	return res
}

func backtrack(startIndex, sum, target int, track, candidates []int, res *[][]int) {
	if sum > target {
			return
	}

	if sum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			*res = append(*res, temp)
			return
	}

	for i := startIndex; i < len(candidates); i++ {
			if i > startIndex && candidates[i] == candidates[i-1] {
					continue
			}

			track = append(track, candidates[i])
			sum += candidates[i]
			backtrack(i+1, sum, target, track, candidates, res)
			sum -= candidates[i]
			track = track[:len(track)-1]
	}
}

func quickSort(nums []int) {
	separateSort(nums, 0, len(nums)-1)
}

func separateSort(nums []int, start, end int) {
	if start > end {
			return
	}

	i := partition(nums, start, end)
	separateSort(nums, start, i-1)
	separateSort(nums, i+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	
	i := start
	for j := start; j < len(nums); j++ {
			if nums[j] < pivot {
					nums[i], nums[j] = nums[j], nums[i]
					i++
			}
	}
	nums[end], nums[i] = nums[i], nums[end]
	return i
}