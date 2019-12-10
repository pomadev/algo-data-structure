package main

import (
	"bufio"
	"fmt"
	"math"
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

	minv, maxd := nextInt(), math.MinInt64
	for i := 1; i < n; i++ {
		r := nextInt()
		if maxd < r-minv {
			maxd = r - minv
		}
		if r < minv {
			minv = r
		}
	}

	fmt.Println(maxd)
}
