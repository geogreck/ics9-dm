package main

import "fmt"

func partition(low, high int, less func(i, j int) bool, swap func(i, j int)) int {
	i := low
	for j := low; j < high; j++ {
		if less(j, high) {
			swap(i, j)
			i++
		}
	}
	swap(i, high)
	return i
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	var qsort_rec func(low, high int)
	qsort_rec = func(low, high int) {
		if low > high {
			return
		}
		q := partition(low, high, less, swap)
		qsort_rec(low, q-1)
		qsort_rec(q+1, high)
	}
	qsort_rec(0, n-1)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	qsort(n, func(i, j int) bool { return arr[i] < arr[j] }, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	for _, x := range arr {
		fmt.Printf("%d ", x)
	}
}
