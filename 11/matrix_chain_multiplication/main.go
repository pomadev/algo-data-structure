package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	p, m := make([]int, n+1), make([][]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&p[i-1], &p[i])
	}

	for i := 1; i <= n; i++ {
		m[i] = make([]int, n+1)
	}
	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt64
			for k := i; k <= j-1; k++ {
				m[i][j] = min(m[i][j], m[i][k]+m[k+1][j]+p[i-1]*p[k]*p[j])
			}
		}
	}

	fmt.Println(m[1][n])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
