// 寻找重复数
package leetcode

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			fast = 0 
			for {
				if slow == fast {
					return fast
				}
				slow = nums[slow]
				fast = nums[fast]
			}
		}
	}
}