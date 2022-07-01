package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type point struct {
	x, y int
}

func dist(a, b point) float64 {
	return float64(math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2)))
}

type edge struct {
	pointA, pointB int
	weight         float64
}

type edgesArr []edge

func (e edgesArr) Len() int {
	return len(e)
}

func (e edgesArr) Less(i, j int) bool {
	return e[i].weight < e[j].weight
}

func (e edgesArr) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func spanningTree(edges []edge, n int) float32 {
	e_ := make([]edge, 0)
	q := make([]Vertex, n)
	for i := 0; i < n; i++ {
		q[i].parent = &q[i]
		q[i].depth = 0
		q[i].i = i
	}
	var sum float64
	for i := 0; i < len(edges) && len(e_) < n-1; i++ {
		u := edges[i].pointA
		v := edges[i].pointB
		if q[u].Find() != q[v].Find() {
			sum += edges[i].weight
			e_ = append(e_, edges[i])
			q[u].Union(&q[v])
		}
	}
	return float32(sum)
}

func kruskal(edges []edge, n int) float32 {
	sort.Sort(edgesArr(edges))
	sum := spanningTree(edges, n)
	return sum
}

func main() {
	var n int
	bufstdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(bufstdin, &n)
	points := make([]point, n)
	for i := range points {
		var x, y int
		fmt.Fscan(bufstdin, &x, &y)
		points[i] = point{x, y}
	}

	edges := make(edgesArr, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(points[i], points[j])})
		}
	}

	fmt.Printf("%.2f\n", kruskal(edges, n))
}
