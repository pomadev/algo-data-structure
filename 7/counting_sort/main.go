package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX = 10000

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

// TODO タイムオーバーしている
func main() {
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		a[i] = nextInt()
	}

	c := make([]int, MAX+1)
	for i := 1; i < n+1; i++ {
		c[a[i]]++
	}
	for i := 1; i < MAX+1; i++ {
		c[i] += c[i-1]
	}

	b := make([]int, n+1)
	for i := n; i >= 1; i-- {
		b[c[a[i]]] = a[i]
		c[a[i]]--
	}
	for i := 1; i < n+1; i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(b[i])
	}
	fmt.Println()
}
