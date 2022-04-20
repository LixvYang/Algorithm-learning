// Package leetcode provides 
package leetcode

var res[][]int
func subsetsWithDup(nums []int)[][]int {
	res=make([][]int,0)
	quickSort(nums)
	dfs([]int{},nums,0)
	return res
}
func dfs(temp, num []int, start int)  {
	tmp:=make([]int,len(temp))
	copy(tmp,temp)

	res=append(res,tmp)
	for i:=start;i<len(num);i++{
		if i>start&&num[i]==num[i-1]{
			continue
		}
		temp=append(temp,num[i])
		dfs(temp,num,i+1)
		temp=temp[:len(temp)-1]
	}
}


func quickSort(nums []int) {
	separateSort(nums, 0, len(nums)-1)
}

func separateSort(nums []int, start, end int) {
	if start > end {
			return
	}

	i := partition(nums, start, end)
	separateSort(nums, start, i-1)
	separateSort(nums, i+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	
	i := start
	for j := start; j < len(nums); j++ {
			if nums[j] < pivot {
					nums[i], nums[j] = nums[j], nums[i]
					i++
			}
	}
	nums[end], nums[i] = nums[i], nums[end]
	return i
}