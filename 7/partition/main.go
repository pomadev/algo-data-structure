package main

import "fmt"

func partition(a []int, p, r int) int {
	x := a[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	q := partition(a, 0, n-1)

	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		if i == q {
			fmt.Print("[")
		}
		fmt.Print(v)
		if i == q {
			fmt.Print("]")
		}
	}
	fmt.Println()
}
