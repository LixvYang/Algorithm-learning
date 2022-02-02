// Package leetcode  provides 
package leetcode 

func partition(s string) [][]string {
	var tmpString []string
	var res [][]string
	backTracking(s,tmpString,0,&res)
	return res
}

func backTracking(s string, tmpString []string, startIndex int, res *[][]string) {
	if startIndex == len(s) {
			t := make([]string, len(tmpString))
			copy(t, tmpString)
			*res = append(*res, t)
	}

	for i := startIndex; i < len(s); i++ {
			if isPartition(s,startIndex,i) {
					tmpString=append(tmpString,s[startIndex:i+1])
			} else {
					continue
			}
			backTracking(s,tmpString,i+1,res)
			tmpString=tmpString[:len(tmpString)-1]
	}
}

func isPartition(s string,startIndex,end int)bool{
	left:=startIndex
	right:=end
	for ;left<right;{
			if s[left]!=s[right]{
					return false
			}
			//移动左右指针
			left++
			right--
	}
	return true
}