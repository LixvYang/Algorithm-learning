package sorting


func BubbleSort(a []int,n int)  {
	if n < 1 {
		return
	}

	for i := 0;i < n;i++ {
		for j := 0;j<n-i-1;j++ {
			if a[j] > a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
}

func InsertSort(a []int,n int)  {
	if n < 1 {
		return
	}
	for i := 1;i < n;i++ {
		for j := 0; j < i;j++ {
			if a[i] < a[j] {
				a[j],a[i] = a[i],a[j]
			}
		}
	}
}

func SelectSort(a []int,n int)  {
	if n < 1 {
		return
	}
	for i := 0;i < n-1;i++ {
		min := i
		for j := i+1;j<n;j++ {
			if a[j]<a[min] {
				a[j],a[min] = a[min],a[j]
			}
		}
	}
}

