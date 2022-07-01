package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vertex struct {
	i, depth int
	parent   *Vertex
}

func MakeSet(x int) *Vertex {
	t := Vertex{}
	t.i = x
	t.depth = 0
	t.parent = &t
	return &t
}

func (x *Vertex) Find() *Vertex {
	var root *Vertex
	if x.parent == x {
		root = x
	} else {
		x.parent = x.parent.Find()
		root = x.parent
	}
	return root
}

func (x *Vertex) Union(y *Vertex) {
	root_x := x.Find()
	root_y := y.Find()
	if root_x.depth < root_y.depth {
		root_x.parent = root_y
	} else {
		root_y.parent = root_x
		if (root_x.depth == root_y.depth) && (root_x != root_y) {
			root_x.depth++
		}
	}
}

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
	fmt.Printf("%d\n%d\n%d\n", this.n, this.m, this.q)
	for _, row := range this.transitionMatrix {
		for _, elem := range row {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
	for _, row := range this.outputMatrix {
		for _, elem := range row {
			fmt.Print(elem, " ")
		}
		fmt.Println()
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

func (this *Automat) VisMealy() {
	fmt.Printf("digraph {\n\trankdir = LR\n")
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.m; j++ {
			fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", i, this.transitionMatrix[i][j], j+97, this.outputMatrix[i][j])
		}
	}
	fmt.Println("}")
}

func (this *Automat) AufenkampHohn() *Automat {
	var m, m1 int
	classes := make([]int, this.n)
	m, classes = this.split1(m, classes)
	for {
		m1, classes = this.split(m1, classes)
		if m == m1 {
			break
		}
		m = m1
	}
	delta := make([]int, this.n)
	phi := make([]int, this.n)
	count := 0
	for i := 0; i < this.n; i++ {
		if classes[i] == i {
			delta[count] = i
			phi[i] = count
			count++
		}
	}
	transitionMatrix := make([][]int, this.n)
	outputMatrix := make([][]string, this.n)
	for i := 0; i < this.n; i++ {
		transitionMatrix[i] = make([]int, this.m)
		outputMatrix[i] = make([]string, this.m)
	}
	newAutomat := Automat{m, this.m, phi[classes[this.q]], transitionMatrix, outputMatrix}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.m; j++ {
			newAutomat.transitionMatrix[i][j] = phi[classes[this.transitionMatrix[delta[i]][j]]]
			/* if i == 4 {
				for _, c := range this.outputMatrix[i] {
					fmt.Printf("%s ", c)
				}
			} */
			newAutomat.outputMatrix[i][j] = this.outputMatrix[delta[i]][j]

		}
	}
	return &newAutomat
}

func (this *Automat) split1(m int, classes []int) (int, []int) {
	m = this.n
	q := make([]Vertex, this.n)
	for i := 0; i < this.n; i++ {
		q[i].parent = &q[i]
		q[i].depth = 0
		q[i].i = i
	}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.n; j++ {
			q1 := q[i]
			q2 := q[j]
			if q1.Find() != q2.Find() {
				eq := true
				for x := 0; x < this.m; x++ {
					if this.outputMatrix[i][x] != this.outputMatrix[j][x] {
						eq = false
						break
					}
				}
				if eq {
					q1.Union(&q2)
					m--
				}
			}
		}
	}
	for i := 0; i < this.n; i++ {
		classes[q[i].i] = q[i].Find().i
	}
	return m, classes
}

func (this *Automat) split(m int, classes []int) (int, []int) {
	m = this.n
	q := make([]Vertex, this.n)
	for i := 0; i < this.n; i++ {
		q[i].parent = &q[i]
		q[i].depth = 0
		q[i].i = i
	}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.n; j++ {
			q1 := q[i]
			q2 := q[j]
			if classes[q1.i] == classes[q2.i] && q1.Find() != q2.Find() {
				eq := true
				for x := 0; x < this.m; x++ {
					w1 := this.transitionMatrix[i][x]
					w2 := this.transitionMatrix[j][x]
					if classes[w1] != classes[w2] {
						eq = false
						break
					}
				}
				if eq {
					q1.Union(&q2)
					m--
				}
			}
		}
	}
	for i := 0; i < this.n; i++ {
		classes[q[i].i] = q[i].Find().i
	}
	return m, classes
}

func (this *Automat) Compare(obj *Automat) bool {
	minCanThis := this.AufenkampHohn().canonize()
	minCanObj := obj.AufenkampHohn().canonize()
	if minCanThis.n != minCanObj.n {
		return false
	}
	if minCanThis.m != minCanObj.m {
		return false
	}
	for i := 0; i < minCanObj.n; i++ {
		for j := 0; j < minCanObj.m; j++ {
			if minCanThis.transitionMatrix[i][j] != minCanObj.transitionMatrix[i][j] {
				return false
			}
			if minCanThis.outputMatrix[i][j] != minCanObj.outputMatrix[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	var a, b Automat
	a.input()
	b.input()
	if a.Compare(&b) {
		fmt.Printf("EQUAL\n")
	} else {
		fmt.Printf("NOT EQUAL\n")
	}
	//b.display()
}
