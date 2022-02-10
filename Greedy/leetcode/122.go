// Package leetcode provides 
package leetcode

func maxProfit(prices []int) int {
	var sum int
	for i := 1; i < len(prices);i++ {
			if prices[i]-prices[i-1]>0 {
					sum+=prices[i]-prices[i-1]
			}
	}
	return sum
}