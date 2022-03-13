package main

import (
	"fmt"
)

func maxmin(a, b int) (max, min int) {
	if a < b {
		return b, a
	} else {
		return a, b
	}
}

func inputLongInt() []int32 {
	var count int
	fmt.Scan(&count)
	num := make([]int32, count)
	for i := range num {
		fmt.Scan(&num[i])
	}
	return num
}

func printLongInt(a []int32) {
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("%d ", a[i])
	}
	fmt.Printf("\n")
}

func clearZero(a []int32) []int32 {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == 0 {
			a = a[0:i]
		} else {
			break
		}
	}
	return a
}

func add(a, b []int32, p int) []int32 {
	max, min := maxmin(len(a), len(b))
	var buf int32 = 0
	ans := make([]int32, max+1)
	for i := 0; i < min; i++ {
		ans[i] = (a[i] + b[i] + buf) % int32(p)
		buf = (a[i] + b[i] + buf) / int32(p)
	}
	if len(a) < len(b) {
		a, b = b, a
	}
	for i := min; i < max; i++ {
		ans[i] = (a[i] + buf) % int32(p)
		buf = (a[i] + buf) / int32(p)
	}
	if buf != 0 {
		ans[len(ans)-1] += buf
	}
	return clearZero(ans)
}

func main() {
	var p int
	fmt.Scan(&p)
	a := inputLongInt()
	b := inputLongInt()
	c := add(a, b, p)
	printLongInt(c)
}
