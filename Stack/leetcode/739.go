// 每日温度
package leetcode

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for k, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			res[top] = k-top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
	}
	return res
}