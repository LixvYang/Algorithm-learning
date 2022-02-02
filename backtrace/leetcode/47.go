// Package leetcode provides 
package leetcode

var res [][]int
func permute(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums,len(nums),[]int{})
	return res
}
func backTrack(nums []int,numsLen int,path []int)  {
	if len(nums)==0{
		p:=make([]int,len(path))
		copy(p,path)
		res = append(res,p)
	}
	used := [21]int{}//跟前一题唯一的区别，同一层不使用重复的数。关于used的思想carl在递增子序列那一题中提到过
	for i:=0;i<numsLen;i++{
		if used[nums[i]+10]==1{
			continue
		}
		cur:=nums[i]
		path = append(path,cur)
		used[nums[i]+10]=1
		nums = append(nums[:i],nums[i+1:]...)
		backTrack(nums,len(nums),path)
		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
		path = path[:len(path)-1]

	}

}