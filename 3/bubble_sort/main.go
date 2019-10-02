package main

import "fmt"

func bubbleSort(a []int) int{
	count := 0
	f := true
	for i := 0; f; i++ {
		f = false
		for j := len(a)-1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				count++
				f = true
			}
		}
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	count := bubbleSort(a)
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(count)
}
