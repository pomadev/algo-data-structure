package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// fmt.Scanでは間に合わない
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextLine())
	return i
}

func main() {
	n := nextInt()

	r := make([]int, n)
	for i := range r {
		r[i] = nextInt()
	}

	maxv, minv := -2000000000, r[0]
	for i := 1; i < n; i++ {
		if r[i] - minv > maxv {
			maxv = r[i]-minv
		}
		if r[i] < minv {
			minv = r[i]
		}
	}

	fmt.Println(maxv)
}
