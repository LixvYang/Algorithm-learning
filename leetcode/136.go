//只出现一次的数字
package leetcode

func singleNumber(nums []int) (ans int) {
	for _, v := range nums {
			ans ^= v
	}
	return ans
}