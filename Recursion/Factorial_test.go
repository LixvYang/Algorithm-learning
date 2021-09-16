package recursion

import (
	"testing"
)

func TestFactorial(t *testing.T)  {
	fac := NewFactorial(10)
	for i := 1;i<5;i++ {
		fac.Factorial(i)
		fac.Print(i)
	}
}