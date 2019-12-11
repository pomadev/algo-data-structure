package main

import "fmt"

func output(a []int, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(a[i])
	}
	fmt.Println()
}

func insertionSort(a []int, n int) {
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v

		output(a, n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	output(a, n)
	insertionSort(a, n)
}
