// 全排列
package leetcode

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, &res, used)
	return res
}

func backtrack(nums, track []int, res *[][]int, used []bool) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	for i := 0;i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack(nums, track, res, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}