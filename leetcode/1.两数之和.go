/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// m[v]index
	for k, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{k, index}
		} else {
			m[v] = k
		}
	}
	return []int{}
}
// @lc code=end

