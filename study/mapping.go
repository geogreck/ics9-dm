package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["alpha"] = 10
	v, ok := a["alpha"]
	fmt.Printf("%d %v\n", v, ok)
	delete(a, "alpha")
	v, ok = a["alpha"]
	fmt.Printf("%d %v\n", v, ok)

	b := map[int]string{
		10: "alpha", 20: "beta", 30: "gamma",
	}
	fmt.Printf("%s\n", b[20])

	c := map[int]bool{10: true, 20: false, 300: true}
	for k, v := range c {
		fmt.Printf("(%d, %v) ", k, v)
	}
}
