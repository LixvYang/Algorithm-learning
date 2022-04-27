package leetcode

func DailyTemperatures(temperatures []int) []int {
	// 单调栈
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					res[top] = i - top
			}

			stack = append(stack, i)
	}
	return res
}