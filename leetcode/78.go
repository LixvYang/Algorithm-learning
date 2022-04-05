// 子集
package leetcode

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(0, nums, []int{}, &res)
	return res
}

func backtrack(startIndex int, nums, track []int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)

	for i := startIndex; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(i+1, nums, track, res)
		track = track[:len(track)-1]
	}

}