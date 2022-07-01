package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	outSig    string
	in, count int
}

type Mealy struct {
	n, m, h          int
	transitionMatrix [][]int
	outputMatrix     [][]int
	inDictionary     []string
	outDictionary    []string
}

type Moore struct {
	nodes         []node
	inDictionary  []string
	outDictionary []string
}

func Lookup(nodes []node, x node) (val node, exists bool) {
	for _, v := range nodes {
		if v.in == x.in && v.outSig == x.outSig {
			return v, true
		}
	}
	return x, false
}

func (this *Mealy) input() {
	bufstdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(bufstdin, &this.m)
	this.inDictionary = make([]string, this.m)
	for i := range this.inDictionary {
		fmt.Fscan(bufstdin, &this.inDictionary[i])
	}

	fmt.Fscan(bufstdin, &this.h)
	this.outDictionary = make([]string, this.h)
	for i := range this.outDictionary {
		fmt.Fscan(bufstdin, &this.outDictionary[i])
	}
	fmt.Fscan(bufstdin, &this.n)
	this.transitionMatrix = make([][]int, this.n)
	for i := range this.transitionMatrix {
		this.transitionMatrix[i] = make([]int, this.m)
		for j := range this.transitionMatrix[i] {
			fmt.Fscan(bufstdin, &this.transitionMatrix[i][j])
		}
	}

	this.outputMatrix = make([][]int, this.n)
	for i := range this.outputMatrix {
		this.outputMatrix[i] = make([]int, this.m)
		for j := range this.outputMatrix[i] {
			fmt.Fscan(bufstdin, &this.outputMatrix[i][j])
		}
	}
}

func (this *Mealy) toMoore() Moore {
	var moore Moore
	moore.nodes = make([]node, 0)

	count := 0
	for i := range this.outputMatrix {
		for j := range this.outputMatrix[i] {
			a := node{fmt.Sprint(this.transitionMatrix[i][j], ",", this.outDictionary[this.outputMatrix[i][j]]), this.transitionMatrix[i][j], count}
			_, exists := Lookup(moore.nodes, a)
			if !exists {
				moore.nodes = append(moore.nodes, a)
				count++
			}
		}
	}
	sort.Slice(moore.nodes, func(i, j int) bool {
		return (moore.nodes[i].in < moore.nodes[j].in || moore.nodes[i].outSig < moore.nodes[j].outSig)
	})

	return moore
}

func (moore Moore) VisMoore(this Mealy) {
	fmt.Print("digraph {\n rankdir = LR\n")
	for _, nod := range moore.nodes {
		fmt.Printf("\t%d\t[label = \"(%s)\"]\n", nod.count, nod.outSig)
		for j := 0; j < this.m; j++ {
			a := node{fmt.Sprint(this.transitionMatrix[nod.in][j], ",", this.outDictionary[this.outputMatrix[nod.in][j]]), this.transitionMatrix[nod.in][j], 0}
			x, _ := Lookup(moore.nodes, a)
			fmt.Printf("\t%v -> %v [label = \"%v\"]\n", nod.count, x.count, this.inDictionary[j])
		}
	}
	fmt.Printf("}\n")
}

func main() {
	var automat Mealy
	automat.input()
	moore := automat.toMoore()
	moore.VisMoore(automat)
}
