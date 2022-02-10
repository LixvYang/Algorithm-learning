// Package leetcode provides 
package leetcode

import (
	"sort"
)
func eraseOverlapIntervals(intervals [][]int) int {
	var flag int
	//先排序
	sort.Slice(intervals,func(i,j int)bool{
			return intervals[i][0]<intervals[j][0]
	})

	for i :=1 ; i<len(intervals);i++ {
			if intervals[i-1][1]>intervals[i][0]{
					flag++
					intervals[i][1]=min(intervals[i-1][1],intervals[i][1])//由于是先排序的，所以，第一位是递增顺序，故只需要将临近两个元素的第二个值最小值更新到该元素的第二个值即可作之后的判断
			}
	}
	return flag
}

func min(a,b int)int {
	if a >b{
			return b
	}
	return a
}