// Package leetcode provides 
package leetcode

func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N)//将数字转为字符串，方便使用下标
	ss := []byte(s)//将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
			return N
	}
	for i:=n-1 ; i>0; i-- {
			if ss[i-1] > ss[i] {//前一个大于后一位,前一位减1，后面的全部置为9
					ss[i-1] -= 1
					for j := i ; j < n; j++ {//后面的全部置为9
							ss[j] = '9'
					}
			} 
	}
	res, _ := strconv.Atoi(string(ss))
	return res 
}