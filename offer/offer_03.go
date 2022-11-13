// 数组中的重复数字
package offer

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}

// func findRepeatNumber(nums []int) int {
// 	numsMap := make(map[int]bool, len(nums))
// 	var b int
// 	for i := 0; i < len(nums); i++ {
// 		if numsMap[nums[i]] {
// 			b = nums[i]
// 			break
// 		}
// 		numsMap[nums[i]] = true
// 	}
// 	return b
// }
