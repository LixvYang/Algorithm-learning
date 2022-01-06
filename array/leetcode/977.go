package leetcode

func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right, k := 0, n-1, n-1
	ans := make([]int, n)

	for left <= right {
			if nums[left] * nums[left] < nums[right] * nums[right] {
					ans[k] = nums[right] * nums[right]
					right--
			} else {
					ans[k] = nums[left] * nums[left]
					left++
			}
			k--
	}
	return ans
}