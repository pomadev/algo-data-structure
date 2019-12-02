package main

import (
	"fmt"
	"math"
)

const INFTY = math.MaxInt64

const (
	WHITE = iota + 1
	GRAY
	BLACK
)

var (
	n int
	g [][]int
)

func prim() int {
	var u, minv int
	d, p, color := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = INFTY
		p[i] = -1
		color[i] = WHITE
	}

	d[0] = 0
	for {
		minv = INFTY
		u = -1
		for i := 0; i < n; i++ {
			if minv > d[i] && color[i] != BLACK {
				u = i
				minv = d[i]
			}
		}
		if u == -1 {
			break
		}
		color[u] = BLACK
		for v := 0; v < n; v++ {
			if color[v] != BLACK && g[u][v] != INFTY {
				if d[v] > g[u][v] {
					d[v] = g[u][v]
					p[v] = u
					color[v] = GRAY
				}
			}
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if p[i] != -1 {
			sum += g[i][p[i]]
		}
	}
	return sum
}

func main() {
	fmt.Scan(&n)

	g = make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&g[i][j])
			if g[i][j] == -1 {
				g[i][j] = INFTY
			}
		}
	}

	fmt.Println(prim())
}
