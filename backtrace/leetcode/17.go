// Package leetcode provides 
package leetcode

func letterCombinations(digits string) []string {
	if len(digits)==0 || len(digits)>4 {
			return nil
	}

	digitsMap := [10]string{
			"", // 0
			"", // 1
			"abc", // 2
			"def", // 3
			"ghi", // 4
			"jkl", // 5
			"mno", // 6
			"pqrs", // 7
			"tuv", // 8
			"wxyz", // 9
	}

	res := make([]string, 0)
	recursion("",digits,0,digitsMap,&res)
	return res
}

func recursion(tempString ,digits string, Index int, digitsMap [10]string, res *[]string) {
	if len(tempString)==len(digits){//终止条件，字符串长度等于digits的长度
			*res=append(*res,tempString)
			return
	}
	tmpK:=digits[Index]-'0' // 将index指向的数字转为int（确定下一个数字）
	letter:=digitsMap[tmpK]// 取数字对应的字符集
	for i:=0;i<len(letter);i++{
			tempString=tempString+string(letter[i])//拼接结果
			recursion(tempString,digits,Index+1,digitsMap,res)
			tempString=tempString[:len(tempString)-1]//回溯
	}
}