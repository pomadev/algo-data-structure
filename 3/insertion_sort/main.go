package main

import "fmt"

func trace(a []int) {
	for i := range a {
		// 左隣にスペース出力
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(a[i])
	}
	fmt.Println()
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		v := a[i]
		j := i-1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
		// 途中結果
		trace(a)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	// 初期状態
	trace(a)
	insertionSort(a)
}
