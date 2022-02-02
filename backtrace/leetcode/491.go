// Package leetcode provides 
package leetcode

func findSubsequences(nums []int) [][]int {
	var subRes []int
	var res [][]int
	backTring(0,nums,subRes,&res)
	return res
}

func backTring(startIndex int, nums []int, subRes []int, res *[][]int) {
	if len(subRes)>1 {
			tmp:=make([]int,len(subRes))
			copy(tmp,subRes)
			*res=append(*res,tmp)
	}

	history := [201]int{}
	for i := startIndex;i<len(nums);i++ {
			if len(subRes)>0&&nums[i]<subRes[len(subRes)-1]||history[nums[i] + 100]==1{
					continue
			}
			history[nums[i] + 100]=1//表示本层该元素使用过了
			subRes=append(subRes,nums[i])
			backTring(i+1,nums,subRes,res)
			subRes=subRes[:len(subRes)-1]

	}
}