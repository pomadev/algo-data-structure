package main

import "fmt"

func search(n, key int, s []int) bool {
	i := 0
	s[n] = key
	for s[i] != key {
		i++
	}
	return i != n
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	var q int
	fmt.Scan(&q)
	t := make([]int, q)
	for i := range t {
		fmt.Scan(&t[i])
	}

	var count int
	for _, v := range t {
		if search(n, v, s) {
			count++
		}
	}
	fmt.Println(count)
}
