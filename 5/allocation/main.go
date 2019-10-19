package main

import "fmt"

var (
	n, k int
	w    []int
)

func check(p int) int {
	i := 0 // 荷物の個数
	for j := 0; j < k; j++ {
		s := 0            // 1台に積む荷物の重さ
		for s+w[i] <= p { // 最大積載量を超えてはいけない
			s += w[i]
			i++
			if i == n { // 全ての荷物を積めた場合
				return n
			}
		}
	}
	return i
}

func solve() int {
	left, right := 0, 100000*10000
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
