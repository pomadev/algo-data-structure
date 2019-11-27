package main

import "fmt"

type Graph [][]int

var (
	dist, que []int
)

func bfs(g Graph) {
	for len(que) != 0 {
		v := que[0]
		que = que[1:]
		for _, nv := range g[v] {
			if dist[nv] != -1 {
				continue
			}
			dist[nv] = dist[v] + 1
			que = append(que, nv)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	g := make(Graph, n)
	dist = make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	for i := 0; i < n; i++ {
		var u, k int
		fmt.Scan(&u, &k)
		u--
		g[u] = make([]int, 0, k)
		for j := 0; j < k; j++ {
			var v int
			fmt.Scan(&v)
			v--
			g[u] = append(g[u], v)
		}
	}
	dist[0] = 0
	que = append(que, 0)
	bfs(g)
	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n", i+1, dist[i])
	}
}
