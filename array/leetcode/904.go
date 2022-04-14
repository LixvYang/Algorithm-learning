// 水果成篮
package leetcode

func totalFruit(fruits []int) int {
	counter := make(map[int]int)
	i, res, n := 0, 0, len(fruits)

	for j := 0; j < n; j++ {
		counter[fruits[j]]++

		for len(count) >= 3 {
			counter[fruits[i]]--

			if counter[fruits[i]] == 0 {
				delete(counter, fruits[i])
			}
			i++
		}

		if res < j-i+1 {
			res = j-i+1
		}
	}
	return res
}