package main

import "fmt"

var c [][]int

func lcs(s1, s2 string) int {
	c = make([][]int, len(s1)+1)
	for i := range c {
		c[i] = make([]int, len(s2)+1)
	}
	maxl := 0
	for i := 1; i < len(c); i++ {
		for j := 1; j < len(c[i]); j++ {
			if s1[i-1] == s2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = max(c[i][j-1], c[i-1][j])
			}
			maxl = max(maxl, c[i][j])
		}
	}
	return maxl
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s1, s2 string
		fmt.Scan(&s1, &s2)
		fmt.Println(lcs(s1, s2))
	}
}
