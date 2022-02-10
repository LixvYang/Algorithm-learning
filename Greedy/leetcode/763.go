// Package leetcode provides 
package leetcode

func partitionLabels(s string) []int {
	var res []int
	var marks [26]int
	size, left, right := len(s),0,0
	for i:=0;i<size;i++ {
			marks[s[i]-'a']=i
	}

	for i:=0;i<size;i++ {
			right = max(right, marks[s[i]-'a'])
			if i==right {
					res = append(res, right-left+1)
					left=i+1
			}
	}
	return res
}
