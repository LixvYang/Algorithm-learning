package recursion

type Fac struct {
	val map[int]int
}

func NewFactorial(n int) *Fac {
	return &Fac{make(map[int]int,n)}
}

func (f *Fac) Factorial(n int) int {
	if f.val[n] != 0 {
		return f.val[n]
	}
	if n <= 1 {
		f.val[n] = 1
		return 1
	} else {
		res := n * f.Factorial(n-1)
		f.val[n] = res
		return res
	}
}

func (f *Fac) Print(n int)  {
	println(f.val[n])
}

