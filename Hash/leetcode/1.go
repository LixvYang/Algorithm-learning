package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, val := range nums {
			if preIndex, ok := m[target - val]; ok {
					return []int{preIndex, index}
			} else {
					m[val] = index
			}
	}
	return []int{}
}