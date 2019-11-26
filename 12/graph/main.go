package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make(map[int]map[int]int, n)

	for i := 0; i < n; i++ {
		a[i] = make(map[int]int, n)
		var u, k int
		fmt.Scan(&u, &k)
		u--
		for j := 0; j < k; j++ {
			var v int
			fmt.Scan(&v)
			v--
			a[u][v] = 1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}
