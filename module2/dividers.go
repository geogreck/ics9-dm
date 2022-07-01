package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func dividers(n int) (divs []int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i*i != n {
				divs = append(divs, n/i)
			}
		}
	}
	sort.Sort(sort.IntSlice(divs))
	return
}

func primereminder(a, b int) bool {
	c := a / b
	arr := dividers(c)
	return len(arr) == 2
}

func main() {
	bufstdin := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(bufstdin, &n)
	if n == 1 {
		fmt.Printf("graph {\n1\n}\n")
		return
	}
	primedivs := dividers(n)
	fmt.Printf("graph {\n")
	for i, x := range primedivs {
		fmt.Printf("\t%d\n", x)
		for j := 0; j < i; j++ {
			if x%primedivs[j] == 0 && primereminder(x, primedivs[j]) {
				fmt.Printf("\t%d -- %d\n", x, primedivs[j])
			}
		}
	}
	fmt.Printf("}\n")
}
