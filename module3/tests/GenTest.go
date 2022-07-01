package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 100
	fmt.Println(n)
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < n; i++ {
		fmt.Println(rand.Intn(n), rand.Intn(n))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			k := rand.Intn(3)
			if k == 0 {
				fmt.Print("x ")
			}
			if k == 1 {
				fmt.Print("y ")
			}
			if k == 2 {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
	fmt.Println(0)
	fmt.Println(20)
}
