package main

import (
	"fmt"
	"sort"
)

const VMAX = 10000

var (
	n, s int
	A    []int
)

func solve() int {
	ans := 0
	B, V := make([]int, n), make([]bool, n)
	for i := 0; i < n; i++ {
		B[i] = A[i]
		V[i] = false
	}
	sort.Ints(B)
	T := make([]int, VMAX+1)
	for i := 0; i < n; i++ {
		T[B[i]] = i
	}
	for i := 0; i < n; i++ {
		if V[i] {
			continue
		}
		cur := i
		S := 0
		m := VMAX
		an := 0
		for {
			V[cur] = true
			an++
			v := A[cur]
			m = min(m, v)
			S += v
			cur = T[v]
			if V[cur] {
				break
			}
		}
		ans += min(S+(an-2)*m, m+S+(an+1)*s)

	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n)
	A = make([]int, n)
	s = VMAX
	for i := range A {
		fmt.Scan(&A[i])
		s = min(s, A[i])
	}

	ans := solve()
	fmt.Println(ans)
}
