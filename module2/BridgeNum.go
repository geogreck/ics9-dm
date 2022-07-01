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

func countBridges(graph []vertex) int { // алгоритм взят с https://e-maxx.ru/algo/bridge_searching
	var dfs func(int, int)
	count := 0
	time := 0
	n := len(graph)
	used := make([]bool, n)
	tin := make([]int, n)
	fup := make([]int, n)
	dfs = func(v, p int) {
		tin[v] = time
		fup[v] = time
		used[v] = true
		time++
		for i := 0; i < len(graph[v].connectedVertex); i++ {
			to := graph[v].connectedVertex[i]
			if to == p {
				continue
			}
			if used[to] {
				fup[v] = min(fup[v], tin[to])
			} else {
				dfs(to, v)
				fup[v] = min(fup[v], fup[to])
				if fup[to] > tin[v] {
					count++
				}

			}
		}
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			dfs(i, -1)
		}
	}
	return count
}

func main() {
	var n, m int
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
	fmt.Println(countBridges(graph))
}
