package leetcode

func TwoSum(nums []int, target int) []int {
	len := len(nums)
	for index, v := range nums {
		for index2 := index + 1; index2 < len; index2++ {
			if v+nums[index2] == target {
				return []int{index, index2}
			}
		}
	}
	return nil
}
