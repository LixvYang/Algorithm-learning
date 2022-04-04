//三数之和
import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
   	n1 := nums[i]
		if n1 > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}	
		
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l<r && n2==nums[l] {
					l++
				}
				for l<r && n3==nums[r] {
					r--
				}
			} else if n1+n2+n3 >0 {
				r--
			} else {
				l++
			}
		}
  }
	return res
}

