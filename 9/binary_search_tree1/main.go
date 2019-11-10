package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	n, _ := strconv.Atoi(nextStr())
	return n
}

type Node struct {
	key                 int
	parent, left, right *Node
}

var (
	n         int
	root, pre *Node
)

func insert(k int) {
	y, x := pre, root
	z := &Node{
		key: k,
	}
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}
}

func preOrder(u *Node) {
	if u == nil {
		return
	}
	fmt.Printf(" %d", u.key)
	preOrder(u.left)
	preOrder(u.right)
}

func inOrder(u *Node) {
	if u == nil {
		return
	}
	inOrder(u.left)
	fmt.Printf(" %d", u.key)
	inOrder(u.right)
}

func main() {
	scanner.Split(bufio.ScanWords)

	n = nextInt()

	for i := 0; i < n; i++ {
		s := nextStr()
		if s == "insert" {
			insert(nextInt())
		} else {
			inOrder(root)
			fmt.Println()
			preOrder(root)
			fmt.Println()
		}
	}
}
