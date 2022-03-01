package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	a := IntSlice{5, 3, 7, 3, 4, 9, 2}
	sort.Sort(a)
	for _, x := range a {
		fmt.Printf("%d ", x)
	}
}
