// Package leetcode provides 
package leetcode

func combinationSum3(k int, n int) [][]int {
	var track []int
	var result [][]int
	backTree(n,k,1,&track,&result)
	return result
}

func backTree(n,k,startIndex int, track *[]int, result *[][]int) {
	if len(*track) == k {
			var sum int
			temp := make([]int, k)
			for k,v:=range *track{
					sum+=v
					temp[k]=v
			}
			if sum > n {
					return 
			}
			if sum == n {
					*result = append(*result, temp)
			}
			return
	}

	for i := startIndex;i <=9; i++ {
			*track=append(*track, i)
			backTree(n,k,i+1,track,result)
			*track=(*track)[:len(*track)-1]
	}
}