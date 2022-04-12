// 每日温度
package leetcode

func dailyTemperatures(num []int) []int {
	ans := make([]int, len(num))
	stack := []int{}

	for i, v := range num {
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i-top
		}

		stack = append(stack, i)
	}
	return ans
}