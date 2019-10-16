package main

import "fmt"

func binarySearch(a []int, key int) bool {
	left := 0
	right := len(a)
	for left < right {
		mid := (left + right) / 2
		if a[mid] == key {
			return true
		}
		if key < a[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	var q int
	fmt.Scan(&q)
	t := make([]int, q)
	for i := range t {
		fmt.Scan(&t[i])
	}

	count := 0
	for _, v := range t {
		if binarySearch(s, v) {
			count++
		}
	}
	fmt.Println(count)
}
