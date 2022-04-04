// 电话号码的字母组合
package leetcode

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) <= 0 {
		return res
	}

	digitsMap := [10]string {
		"", //0
		"", //1
		"abc", //2
		"def", //3
		"ghi", //4
		"jkl", // 5
     "mno", // 6
     "pqrs", // 7
     "tuv", // 8
     "wxyz", // 9
	}
	
	backtrack(0, &res, digits,"", digitsMap)
	return res
}

func backtrack(Index int, res *[]string, digits, tmpString string, digitsMap [10]string) {
	if len(tmpString) == len(digits) {
		*res = append(*res, tmpString)
		return
	}

	tmpIndex := digits[Index] - '0'
	letter := digitsMap[tmpIndex]
	for i := 0; i < len(letter); i++ {
		tempString := tmpString + string(letter[i])
		backtrack(Index+1, res, digits, tempString, digitsMap)
		tempString = tempString[:len(tempString)-1]
	}
}