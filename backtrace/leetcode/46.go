// Package leetcode provides 
package leetcode

var res [][]int
func permute(nums []int) [][]int {
    res = [][]int{}
    backTrack(nums,len(nums),[]int{})
	return res
}

func backTrack(nums []int, numsLen int, path []int) {
    if len(nums)==0{
		p:=make([]int,len(path))
		copy(p,path)
		res = append(res,p)
	}

    for i:=0;i<numsLen;i++ {
        cur := nums[i]
        path = append(path, cur)
        nums = append(nums[:i], nums[i+1:]...)
        backTrack(nums,len(nums),path)
        nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
        path = path[:len(path)-1]
    }
}