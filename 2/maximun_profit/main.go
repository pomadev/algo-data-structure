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

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt()
	}

	mint, maxd := 0, r[1]-r[0]
	for i := 1; i < n; i++ {
		if maxd < r[i]-r[mint] {
			maxd = r[i] - r[mint]
		}
		if r[i] < r[mint] {
			mint = i
		}
	}

	fmt.Println(maxd)
}
