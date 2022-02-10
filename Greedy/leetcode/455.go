// Package leetcode provides 
package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(s)
	sort.Ints(g)	

	child := 0
	for sIdx := 0; child<len(g)&&sIdx<len(s);sIdx++ {
		if g[child]<=s[sIdx] {
			child++
		}
	}
	return child
}