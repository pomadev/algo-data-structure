package main

import (
	"fmt"
	"io"
	"strconv"
)

var (
	top int
	s   [1000]int
)

func push(x int) {
	top++
	s[top] = x
}

func pop() int {
	top--
	return s[top+1]
}

func main() {
	var (
		a, b int
		s    string
	)
	top = 0
	for {
		_, err := fmt.Scan(&s)
		if err == io.EOF {
			break
		}
		if s == "+" {
			a = pop()
			b = pop()
			push(a + b)
		} else if s == "-" {
			b = pop()
			a = pop()
			push(a - b)
		} else if s == "*" {
			a = pop()
			b = pop()
			push(a * b)
		} else {
			n, _ := strconv.Atoi(s)
			push(n)
		}
	}

	fmt.Println(pop())
}
