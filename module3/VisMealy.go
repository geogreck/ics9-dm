package main

import "fmt"

type Automat struct {
	n, m, q          int
	transitionMatrix [][]int
	outputMatrix     [][]string
}

func (this *Automat) input() {
	fmt.Scan(&this.n)
	fmt.Scan(&this.m)
	fmt.Scan(&this.q)
	this.transitionMatrix = make([][]int, this.n)
	this.outputMatrix = make([][]string, this.n)
	for i := 0; i < this.n; i++ {
		this.transitionMatrix[i] = make([]int, this.m)
		this.outputMatrix[i] = make([]string, this.m)
	}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.m; j++ {
			fmt.Scan(&this.transitionMatrix[i][j])
		}
	}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.m; j++ {
			fmt.Scan(&this.outputMatrix[i][j])
		}
	}
}

func (this *Automat) VisMealy() {
	fmt.Printf("digraph {\n\trankdir = LR\n")
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.m; j++ {
			fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", i, this.transitionMatrix[i][j], j+97, this.outputMatrix[i][j])
		}
	}
	fmt.Println("}")
}

func main() {
	var automat Automat
	automat.input()
	automat.VisMealy()

}
