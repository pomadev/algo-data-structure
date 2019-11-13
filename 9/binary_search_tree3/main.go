package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	n, _ := strconv.Atoi(nextStr())
	return n
}

type Node struct {
	key     int
	p, l, r *Node
}

var (
	r *Node
	b bytes.Buffer
)

func insert(k int) {
	z := &Node{
		key: k,
	}
	x := r
	var y *Node
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.l
		} else {
			x = x.r
		}
	}
	z.p = y
	if r == nil {
		r = z
	} else {
		if z.key < y.key {
			y.l = z
		} else {
			y.r = z
		}
	}
}

func find(u *Node, k int) *Node {
	for u != nil && u.key != k {
		if k < u.key {
			u = u.l
		} else {
			u = u.r
		}
	}
	return u
}

func preOrder(u *Node) {
	if u == nil {
		return
	}
	b.Write([]byte(fmt.Sprintf(" %d", u.key)))
	preOrder(u.l)
	preOrder(u.r)
}

func inOrder(u *Node) {
	if u == nil {
		return
	}
	inOrder(u.l)
	b.Write([]byte(fmt.Sprintf(" %d", u.key)))
	inOrder(u.r)
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	for i := 0; i < n; i++ {
		s := nextStr()
		switch s {
		case "insert":
			insert(nextInt())
		case "find":
			if find(r, nextInt()) != nil {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "delete":
			nextInt()
		case "print":
			inOrder(r)
			b.WriteTo(os.Stdout)
			fmt.Println()
			preOrder(r)
			b.WriteTo(os.Stdout)
			fmt.Println()
		}
	}
}
