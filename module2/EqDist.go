package main

import (
	"bufio"
	"fmt"
	"os"
)

type vertex struct {
	connectedVertex []int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BFS(graph []vertex, v int) []int {
	n := len(graph)
	ans := make([]int, n)
	queue := make(chan int, len(graph))
	queue <- v
	visited := make([]bool, n)
	for len(queue) != 0 {
		w := <-queue
		visited[w] = true
		for _, i := range graph[w].connectedVertex {
			if visited[i] == true {
				continue
			}
			visited[i] = true
			ans[i] = ans[w] + 1
			queue <- i
		}
	}
	return ans
}

func countEqDistances(graph []vertex, roots []int) []int {
	n := len(graph)
	k := len(roots)
	distances := make([][]int, k)
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		distances[i] = BFS(graph, roots[i])
	}
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < k-1; j++ {
			if distances[j][i] != distances[j+1][i] || distances[j][i] == 0 {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	var n, m, k int
	bufstdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(bufstdin, &n)
	fmt.Fscan(bufstdin, &m)
	graph := make([]vertex, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(bufstdin, &a, &b)
		graph[a].connectedVertex = append(graph[a].connectedVertex, b)
		graph[b].connectedVertex = append(graph[b].connectedVertex, a)
	}
	fmt.Fscan(bufstdin, &k)
	roots := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(bufstdin, &roots[i])
	}
	f := bufio.NewWriter(os.Stdout)
	ans := countEqDistances(graph, roots)
	if len(ans) == 0 {
		fmt.Fprintf(f, "-\n")
	} else {
		for _, x := range ans {
			fmt.Fprintf(f, "%d ", x)
		}
	}
	fmt.Fprintf(f, "\n")
	f.Flush()
}
