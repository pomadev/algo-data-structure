package main

import "fmt"

var f = make([]int, 45)

func init() {
	f[0], f[1] = 1, 1
	for i := 2; i <= 44; i++ {
		f[i] = f[i-1] + f[i-2]
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(f[n])
}
