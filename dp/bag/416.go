package bag

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum % 2 != 0 {
		return false
	}
	target := sum/2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		
	}
	
}