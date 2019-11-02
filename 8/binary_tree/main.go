package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const NIL = -1

type Node struct {
	parent, left, right int
}

var (
	n    int
	D, H []int
	T    []Node
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func print(u int) {
	fmt.Printf("node %d: parent = %d, ", u, T[u].parent)
	var s int
	if T[u].parent == NIL {
		s = NIL
	} else {
		s1, s2 := T[T[u].parent].left, T[T[u].parent].right
		if s1 != u {
			s = s1
		} else {
			s = s2
		}
	}
	fmt.Printf("sibling = %d, ", s)
	var c int
	if T[u].left != NIL {
		c++
	}
	if T[u].right != NIL {
		c++
	}
	fmt.Printf("degree = %d, depth = %d, height = %d, ", c, D[u], H[u])
	if T[u].parent == NIL {
		fmt.Println("root")
	} else if T[u].left == NIL && T[u].right == NIL {
		fmt.Println("leaf")
	} else {
		fmt.Println("internal node")
	}
}

func setDepth(u, p int) {
	D[u] = p
	if T[u].right != NIL {
		setDepth(T[u].right, p+1)
	}
	if T[u].left != NIL {
		setDepth(T[u].left, p+1)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func setHeight(u int) int {
	var h1, h2 int
	if T[u].left != NIL {
		h1 = setHeight(T[u].left) + 1
	}
	if T[u].right != NIL {
		h2 = setHeight(T[u].right) + 1
	}
	H[u] = max(h1, h2)
	return H[u]
}

func main() {
	scanner.Split(bufio.ScanWords)

	n = nextInt()
	T = make([]Node, n)
	for i := range T {
		T[i].parent, T[i].left, T[i].right = NIL, NIL, NIL
	}

	for i := 0; i < n; i++ {
		v := nextInt()
		T[v].left, T[v].right = nextInt(), nextInt()
		if T[v].left != -1 {
			T[T[v].left].parent = v
		}
		if T[v].right != -1 {
			T[T[v].right].parent = v
		}
	}

	D, H = make([]int, n), make([]int, n)
	var r int
	for i := range T {
		if T[i].parent == NIL {
			r = i
		}
	}
	setDepth(r, 0)
	setHeight(r)

	for i := 0; i < n; i++ {
		print(i)
	}
}
