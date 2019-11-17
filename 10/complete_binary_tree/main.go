package main

import "fmt"

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func main() {
	var h int
	fmt.Scan(&h)

	a := make([]int, h+1) // 1オリジンにするため+1
	for i := 1; i < h+1; i++ {
		fmt.Scan(&a[i])
	}

	for i := 1; i < h+1; i++ {
		fmt.Printf("node %d: key = %d, ", i, a[i])
		if 1 <= parent(i) && parent(i) <= h {
			fmt.Printf("parent key = %d, ", a[parent(i)])
		}
		if left(i) <= h {
			fmt.Printf("left key = %d, ", a[left(i)])
		}
		if right(i) <= h {
			fmt.Printf("right key = %d, ", a[right(i)])
		}
		fmt.Println()
	}
}
