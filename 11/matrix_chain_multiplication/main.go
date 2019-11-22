package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	r, c := make([]int, 6), make([]int, 6)
	for i := 0; i < n; i++ {
		fmt.Scan(&r[i], &c[i])
	}
}
