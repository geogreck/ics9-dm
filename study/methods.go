package main

import "fmt"

type IntSlice []int

func (slice IntSlice) Len() int { return len(slice) }

func (slice *IntSlice) Append(x int) {
	*slice = append(*slice, x)
}

func main() {
	a := make(IntSlice, 0, 5)
	for i := 0; i < 10; i++ {
		a.Append(i)
	}
	for _, x := range a {
		fmt.Printf("%d ", x)
	}
}
