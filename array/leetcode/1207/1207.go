package leetcode

func uniqueOccurrences(arr []int) bool {
	count := make([]int, 2002)
	// wg := sync.WaitGroup{}
	// for i := 0; i < len(arr); i++ {
	//     wg.Add(1)
	//     go func(i int) {
	//         count[arr[i]+1000]++
	//         wg.Done()
	//     }(i)
	// }
	// wg.Wait()

	for i := 0; i < len(arr); i++ {
		count[arr[i]+1000]++
	}

	fre := make([]bool, 1002)
	for i := 0; i <= 2000; i++ {
		if count[i] > 0 {
			if !fre[count[i]] {
				fre[count[i]] = true
			} else {
				return false
			}
		}
	}
	return true
}
