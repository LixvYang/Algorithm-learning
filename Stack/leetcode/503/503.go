package leetcode

func NextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	stack := make([]int, 0)

	for i := 0; i < 2*n; i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			res[top] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}
