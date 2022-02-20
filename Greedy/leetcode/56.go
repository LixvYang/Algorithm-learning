// Package leetcode provides 
package leetcode

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals,func(i,j int)bool{
			return intervals[i][0]<intervals[j][0]
	})

	//再弄重复的
	for i:=0;i<len(intervals)-1;i++{
			if intervals[i][1]>=intervals[i+1][0]{
					intervals[i][1]=max(intervals[i][1],intervals[i+1][1])//赋值最大值
					intervals=append(intervals[:i+1],intervals[i+2:]...)
					i--
			}
	}
	return intervals
}