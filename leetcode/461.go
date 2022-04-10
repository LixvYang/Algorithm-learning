// 汉明距离
func hammingDistance(x int, y int) int {
	z := x ^ y
	fmt.Println(z)
	var res int
	for z > 0 {
			z -= lowbit(z)
			res += 1
	}
	return res
}

func lowbit(x int) int {
	return x & (-x)
}