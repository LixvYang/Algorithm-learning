package leetcode

type NumArray struct { // 维护了preSum数组和nums数组的长度
	preSum  []int
	numsLen int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 { // nums数组没有元素 就不能preSum[0] = nums[0] 这么写
		return NumArray{
			preSum:  []int{},
			numsLen: 0,
		}
	}
	nA := NumArray{ // nums数组有元素，创建NumArray
		preSum:  make([]int, len(nums)), // 每个元素对应一个前缀和，所以preSum和nums等长
		numsLen: len(nums),
	}
	nA.preSum[0] = nums[0] // 预置边界情况
	for i := 1; i < len(nums); i++ {
		nA.preSum[i] = nA.preSum[i-1] + nums[i] // 套用递推式 求出preSum[i]
	}
	return nA
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 { // 此时preSum[i-1]应该为0，从0到j的求和，应该返回preSum[j]
		if this.numsLen == 0 { // 但如果nums根本没有长度，直接返回0
			return 0
		}
		return this.preSum[j]
	}
	return this.preSum[j] - this.preSum[i-1] // 非i=0情况，套用通式
}