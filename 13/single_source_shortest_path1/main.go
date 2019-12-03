package main

import (
	"fmt"
	"math"
)

const (
	WHITE = iota + 1
	GRAY
	BLACK
)

var (
	n int
	g map[int]map[int]int
)

func dijkstra(s int) {
	color, d, p := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = WHITE
		d[i] = math.MaxInt64
	}
	d[s] = 0
	p[s] = -1

	for {
		var u int
		mincost := math.MaxInt64
		for i := 0; i < n; i++ {
			if color[i] != BLACK && d[i] < mincost {
				mincost = d[i]
				u = i
			}
		}
		if mincost == math.MaxInt64 {
			break
		}
		color[u] = BLACK
		for v := 0; v < n; v++ {
			if color[v] == BLACK {
				continue
			}
			if _, ok := g[u][v]; !ok {
				continue
			}
			if d[u]+g[u][v] < d[v] {
				d[v] = d[u] + g[u][v]
				p[v] = u
				color[v] = GRAY
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n", i, d[i])
	}
}

func main() {
	fmt.Scan(&n)
	g = make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		var u, k int
		fmt.Scan(&u, &k)
		g[u] = make(map[int]int)
		for j := 0; j < k; j++ {
			var v, c int
			fmt.Scan(&v, &c)
			g[u][v] = c
		}
	}
	dijkstra(0)
}
