// Package leetcode provides 
package leetcode

func lemonadeChange(bills []int) bool {
	left := [2]int{0,0}
	if bills[0]!=5 {
			return false
	}

	for i:=0;i< len(bills);i++ {
			if bills[i]==5 {
					left[0]++
			}
			if bills[i] == 10 {
					left[1]++
			}
			tmp := bills[i]-5
			if tmp==5 {
					if left[0]>0 {
							left[0]--
					} else {
							return false
					}
			}
			if tmp == 15 {
					if left[1]>0&&left[0]>0 {
							left[0]--
							left[1]--
					} else if left[0]>2 && left[1]==0 {
							left[0]-=3
					} else {
							return false
					}
			}
	}
	return true
}