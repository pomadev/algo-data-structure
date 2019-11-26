package main

import "fmt"

type Graph [][]int

var (
	seen []bool
	d, f []int
	ptr  int
)

func dfs(g Graph, v int) {
	d[v] = ptr
	ptr++

	seen[v] = true
	for _, nv := range g[v] {
		if seen[nv] {
			continue
		}
		dfs(g, nv)
	}

	f[v] = ptr
	ptr++
}

func main() {
	var n int
	fmt.Scan(&n)

	g := make(Graph, n)

	for i := 0; i < n; i++ {
		var u, k int
		fmt.Scan(&u, &k)
		u-- // 0オリジンにする
		g[u] = make([]int, 0, k)
		for j := 0; j < k; j++ {
			var v int
			fmt.Scan(&v)
			v--
			g[u] = append(g[u], v)
		}
	}

	seen, d, f = make([]bool, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		if !seen[i] {
			dfs(g, i)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d %d %d\n", i+1, d[i]+1, f[i]+1)
	}
}
