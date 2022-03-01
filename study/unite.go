package main

import "fmt"

func unite(in1, in2, out chan int) {
	open1, open2 := true, true
	for open1 || open2 {
		var x int
		select {
		case x, open1 = <-in1:
			if open1 {
				out <- x
			}
		case x, open2 = <-in2:
			if open2 {
				out <- x
			}
		}
	}
	close(out)
}

func generate(start, finish int, out chan int) {
	for i := start; i <= finish; i++ {
		out <- i
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch := make(chan int)
	go generate(1, 10, ch1)
	go generate(50, 55, ch2)
	go unite(ch1, ch2, ch)
	for x := range ch {
		fmt.Printf("%d ", x)
	}
}
