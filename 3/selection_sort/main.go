package main

import "fmt"

func selectionSort(a []int) int {
	count := 0
	for i := 0; i < len(a)-1; i++ {
		minj := i
		for j := i; j < len(a); j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		a[i], a[minj] = a[minj], a[i]
		if i != minj {
			count++
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

	count := selectionSort(a)
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(count)
}
