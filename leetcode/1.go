// 两数之和
package leetcode 

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
			if preIndex, ok := m[target-value]; ok {
					return []int{preIndex, index}
			} else {
					m[value] = index
			}
	}
	return []int{}
}

func twosum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}