package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	m, cnt int
	g      []int
)

func insertion(a []int, g int) {
	for i := g; i < len(a); i++ {
		v := a[i]
		j := i - g
		for j >= 0 && v < a[j] {
			a[j+g] = a[j]
			j -= g
			cnt++
		}
		a[j+g] = v
	}
}

func shellSort(a []int) {
	// 数列 {1, 4, 13, 40, 121, ...} を生成
	for h := 1; ; {
		if h > len(a) {
			break
		}
		g = append(g, h)
		h = 3*h + 1
	}

	m = len(g)
	for i := m - 1; i >= 0; i-- {
		insertion(a, g[i])
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	s := nextLine()
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}

	shellSort(a)

	fmt.Println(m)
	for i := m - 1; i >= 0; i-- {
		fmt.Print(g[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	fmt.Println(cnt)
	for _, v := range a {
		fmt.Println(v)
	}
}
