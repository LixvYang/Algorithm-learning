// Package leetcode provides 
package leetcode

func candy(ratings []int) int {
	need:=make([]int,len(ratings))
	sum:=0
	 
	 for i:=0;i<len(ratings);i++{
			 need[i]=1
	 }

	for i:=0;i<len(ratings)-1;i++ {
			if ratings[i]<ratings[i+1] {
					need[i+1]=need[i]+1
			}
	}

	for i:=len(ratings)-1;i>0;i-- {
			if ratings[i-1]>ratings[i] {
					need[i-1]= max(need[i]+1, need[i-1])
			}
	}
	for i:=0;i<len(ratings);i++ {
			sum+=need[i]
	}
	return sum
}
