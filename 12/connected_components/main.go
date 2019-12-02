package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

type Graph [][]int

var (
	seen []bool
	f    []int
)

func dfs(g Graph, v, c int) {
	seen[v] = true
	f[v] = c

	for _, nv := range g[v] {
		if seen[nv] {
			continue
		}
		dfs(g, nv, c)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()

	g := make(Graph, n)
	for i := 0; i < m; i++ {
		s, t := nextInt(), nextInt()
		g[s] = append(g[s], t)
		g[t] = append(g[t], s)
	}

	seen, f = make([]bool, n), make([]int, n)
	var c int
	for i := 0; i < n; i++ {
		if !seen[i] {
			dfs(g, i, c)
		}
		c++
	}

	q := nextInt()
	for i := 0; i < q; i++ {
		s, t := nextInt(), nextInt()
		if f[s] == f[t] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
