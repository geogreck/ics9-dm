package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type edge struct {
	x, y, length int
}

type vertex struct {
	edges      []edge
	index, key int
}

type PriorityQueue []*vertex

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].key < pq[j].key //???? знак сравнения
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*vertex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *vertex, key int) {
	item.key = key
	heap.Fix(pq, item.index)
}

func prim(graph []vertex) int {
	n := len(graph)
	for i := 0; i < n; i++ {
		graph[i].index = -1
	}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	sum := 0
	pos := &graph[0]
	for {
		pos.index = -2
		for _, e := range pos.edges {
			u := &graph[e.y]
			if u.index == -1 {
				u.key = e.length
				heap.Push(&pq, u)
			} else if u.index != -2 && e.length < u.key {
				pq.update(u, e.length)
			}
		}
		if len(pq) == 0 {
			break
		}
		pos = heap.Pop(&pq).(*vertex)
		sum += pos.key
	}
	return sum
}

func main() {
	var n, m int
	bufstdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(bufstdin, &n)
	fmt.Fscan(bufstdin, &m)
	graph := make([]vertex, n)
	for i := 0; i < m; i++ {
		var x, y, len int
		fmt.Fscan(bufstdin, &x, &y, &len)
		e := edge{x, y, len}
		graph[x].edges = append(graph[x].edges, e)
		e = edge{y, x, len}
		graph[y].edges = append(graph[y].edges, e)
	}

	fmt.Println(prim(graph))
}
