package main

import (
	"bufio"
	"fmt"
	"os"
)

type Automat struct {
	n, m, q          int
	transitionMatrix [][]int
	outputMatrix     [][]string
}

func (this *Automat) input() {
	bufstdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(bufstdin, &this.n)
	fmt.Fscan(bufstdin, &this.m)
	fmt.Fscan(bufstdin, &this.q)
	this.transitionMatrix = make([][]int, this.n)
	this.outputMatrix = make([][]string, this.n)
	for i := 0; i < this.n; i++ {
		this.transitionMatrix[i] = make([]int, this.m)
		this.outputMatrix[i] = make([]string, this.m)
	}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.m; j++ {
			fmt.Fscan(bufstdin, &this.transitionMatrix[i][j])
		}
	}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.m; j++ {
			fmt.Fscan(bufstdin, &this.outputMatrix[i][j])
		}
	}
}

func (this *Automat) display() {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fprintf(f, "%d\n%d\n%d\n", this.n, this.m, this.q)
	for _, row := range this.transitionMatrix {
		for _, elem := range row {
			fmt.Fprintf(f, "%d ", elem)
		}
		fmt.Fprintf(f, "\n")
	}
	for _, row := range this.outputMatrix {
		for _, elem := range row {
			fmt.Fprintf(f, "%s ", elem)
		}
		fmt.Fprintf(f, "\n")
	}
}

func (this *Automat) canonize() *Automat {
	transitionMatrix := make([][]int, this.n)
	outputMatrix := make([][]string, this.n)
	for i := 0; i < this.n; i++ {
		transitionMatrix[i] = make([]int, this.m)
		outputMatrix[i] = make([]string, this.m)
	}
	newAutomat := Automat{this.n, this.m, 0, transitionMatrix, outputMatrix}
	used := this.DFS()
	for i := 0; i < this.n; i++ {
		if used[i] != -1 {
			newAutomat.outputMatrix[used[i]] = this.outputMatrix[i]
			for j := 0; j < this.m; j++ {
				newAutomat.transitionMatrix[used[i]][j] = used[this.transitionMatrix[i][j]]
			}
		}
	}
	return &newAutomat
}

func (this *Automat) DFS() []int {
	visited := make([]bool, this.n)
	used := make([]int, this.n)
	count := 0
	for i := 0; i < this.n; i++ {
		used[i] = -1
	}
	var VisitVertex func(int)
	VisitVertex = func(begin int) {
		used[begin] = count
		count++
		visited[begin] = true
		for i := 0; i < this.m; i++ {
			if used[this.transitionMatrix[begin][i]] == -1 {
				VisitVertex(this.transitionMatrix[begin][i])
			}
		}
	}
	VisitVertex(this.q)
	return used
}

func main() {
	var automat Automat
	automat.input()
	a := automat.canonize()
	a.display()
}
