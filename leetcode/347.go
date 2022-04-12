// 前K个高频元素
package leetcode

func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	nums_map := map[int]int{}
	for _, num := range nums {
			nums_map[num]++
	}
	fmt.Println(nums_map)
	for key, _ := range nums_map {
			ans = append(ans, key)
	}
	fmt.Println(ans)

	sort.Slice(ans, func(a, b int) bool {
			return nums_map[ans[a]] > nums_map[ans[b]]
	})
	return ans[:k]
}