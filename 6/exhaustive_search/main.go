package main

import "fmt"

var (
	n int
	a []int
)

// i番目以降の要素を用いて、mを作ることができるかどうか
func solve(i, m int) bool {
	if m == 0 {
		return true
	}
	if i >= n {
		return false
	}
	res := solve(i+1, m) || solve(i+1, m-a[i])
	return res
}

func main() {
	fmt.Scan(&n)
	a = make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var m int
		fmt.Scan(&m)
		if solve(0, m) {
			fmt.Println("yes")
			continue
		}
		fmt.Println("no")
	}
}
