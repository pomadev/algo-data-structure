package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	p, l, r int
}

var (
	n int
	D []int
	T []Node
)

func print(u int) {
	fmt.Printf("node %d: parent = %d, depth = %d, ", u, T[u].p, D[u])
	if T[u].p == -1 {
		fmt.Print("root, ")
	} else if T[u].l == -1 {
		fmt.Print("leaf, ")
	} else {
		fmt.Print("internal node, ")
	}

	fmt.Print("[")
	for i, c := 0, T[u].l; c != -1; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(c)
		c = T[c].r
	}
	fmt.Println("]")
}

func rec(u, p int) {
	D[u] = p
	if T[u].r != -1 {
		rec(T[u].r, p)
	}
	if T[u].l != -1 {
		rec(T[u].l, p+1)
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	scanner.Split(bufio.ScanWords)

	n = nextInt()
	T, D = make([]Node, n), make([]int, n)
	for i := range T {
		T[i].p, T[i].l, T[i].r = -1, -1, -1
	}

	for i := 0; i < n; i++ {
		var l int
		v, d := nextInt(), nextInt()
		for j := 0; j < d; j++ {
			c := nextInt()
			if j == 0 {
				T[v].l = c
			} else {
				T[l].r = c
			}
			l = c
			T[c].p = v
		}
	}

	var r int
	for i := 0; i < n; i++ {
		if T[i].p == -1 {
			r = i
			break
		}
	}

	rec(r, 0)

	for i := 0; i < n; i++ {
		print(i)
	}
}
