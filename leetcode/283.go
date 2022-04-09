// 移动零
package leetcode

func moveZeroes(nums []int){
	left,right,n := 0,0,len(nums)
	for right < n{
			if nums[right] != 0{
					nums[right],nums[left] = nums[left],nums[right]
					left++
			}
			right++
	}
}

