package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]string, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(a)
}
