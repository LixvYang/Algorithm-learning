package leetcode

func removeElement(nums []int, val int) int {
	res := 0 
	for i := 0; i < len(nums); i++ {
			if nums[i] != val {
					nums[res] = nums[i]
					res++
			}
	}
	return res
}

func removeElement2(nums []int, val int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
			if nums[i] == val {
					for j := i+1; j < length; j++ {
							nums[j-1] = nums[j]
					}
					i--
					length--
			}
	}
	return length
}