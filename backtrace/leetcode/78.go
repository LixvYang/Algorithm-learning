// Package leetcode provides 
package leetcode

var res [][]int
func subsets(nums []int) [][]int {
    res = make([][]int, 0)
    sort.Ints(nums)
    Dfs([]int{}, nums, 0)
    return res
}

func Dfs(temp, nums []int, start int) {
    tmp := make([]int, len(temp))
	copy(tmp, temp)
    res = append(res, tmp)

    for i := start;i<len(nums);i++ {
        temp = append(temp, nums[i])
        Dfs(temp, nums, i+1)
        temp = temp[:len(temp)-1]
    }
}

func subsets(nums []int) [][]int {
    res := [][]int{}
    backtrack(0, nums, []int{}, &res)
    return res
}

func backtrack(startIndex int, nums []int, track []int, res *[][]int) {
    temp:=make([]int, len(track))
    copy(temp, track)
    *res = append(*res, temp)

    for i:=startIndex;i<len(nums);i++ {
        track = append(track, nums[i])
        backtrack(i+1, nums, track, res)
        track = track[:len(track)-1]
    } 
}