package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

type Node struct {
	parent, left, right int
}

var (
	n int
	T []Node
)

const NIL = -1

func preParse(u int) {
	if u == NIL {
		return
	}
	fmt.Print(" ", u)
	preParse(T[u].left)
	preParse(T[u].right)
}

func inParse(u int) {
	if u == NIL {
		return
	}
	inParse(T[u].left)
	fmt.Print(" ", u)
	inParse(T[u].right)
}

func postParse(u int) {
	if u == NIL {
		return
	}
	postParse(T[u].left)
	postParse(T[u].right)
	fmt.Print(" ", u)
}

func main() {
	scanner.Split(bufio.ScanWords)

	n = nextInt()
	T = make([]Node, n)
	for i := range T {
		T[i].parent, T[i].left, T[i].right = -1, -1, -1
	}
	for i := 0; i < n; i++ {
		v := nextInt()
		T[v].left, T[v].right = nextInt(), nextInt()
		if T[v].left != NIL {
			T[T[v].left].parent = v
		}
		if T[v].right != NIL {
			T[T[v].right].parent = v
		}
	}

	var r int
	for i := range T {
		if T[i].parent == NIL {
			r = i
		}
	}

	fmt.Println("Preorder")
	preParse(r)
	fmt.Println()

	fmt.Println("Inorder")
	inParse(r)
	fmt.Println()

	fmt.Println("Postorder")
	postParse(r)
	fmt.Println()
}
