// 除以自身以外的最大乘积
package leetcode

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for k, _ := range res {
		res[k] = 1
	}

	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		res[i] *= left
		left *= nums[i]
		res[n-1-i] *= right
		right *= nums[n-1-i]
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	res := make([]int, len(nums))
	for k, _ := range res {
		res[k] = 1
	}

	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		res[i] *= left
		left *= nums[i]
	}

	for j := len(nums)-1; j >= 0; j-- {
		res[j] *= right
		right *= nums[j]
	}
	fmt.Println(res)
	return res
}