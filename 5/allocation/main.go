package main

import "fmt"

var (
	n, k int
	w    []int
)

func check(p int) int {
	i := 0
	for j := 0; j < k; j++ {
		s := 0
		for s+w[i] <= p {
			s += w[i]
			i++
			if i == n {
				return n
			}
		}
	}
	return i
}

func solve() int {
	left, right := 0, 100000*100000
	for left+1 < right {
		mid := (left + right) / 2
		v := check(mid)
		if v >= n {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func main() {
	fmt.Scan(&n, &k)

	w = make([]int, n)
	for i := range w {
		fmt.Scan(&w[i])
	}

	fmt.Println(solve())
}
