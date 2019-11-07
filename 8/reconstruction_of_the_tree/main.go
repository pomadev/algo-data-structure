package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	n, pos        int
	pre, in, post []int
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func find(a []int, n int) int {
	for i, v := range a {
		if v == n {
			return i
		}
	}
	return -1
}

func reconstruction(l, r int) {
	if l >= r {
		return
	}
	root := pre[pos]
	pos++
	m := find(in, root)

	reconstruction(l, m)
	reconstruction(m+1, r)

	post = append(post, root)
}

func main() {
	scanner.Split(bufio.ScanWords)

	n = nextInt()

	pre = make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = nextInt()
	}
	in = make([]int, n)
	for i := 0; i < n; i++ {
		in[i] = nextInt()
	}

	reconstruction(0, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(post[i])
	}
	fmt.Println()
}
