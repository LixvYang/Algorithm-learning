// 根据身高重建队列
package leetcode

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	result := [][]int{}
	for _, info := range people {
		result = append(result, info)
		copy(result[info[1]+1:], result[info[1]:])
		result[info[1]] = info
	}
	return result
}