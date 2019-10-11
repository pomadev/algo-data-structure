package main

import (
	"fmt"
)

type Stack struct {
	s   []int
	top int
}

func (s *Stack) Push(x int) {
	s.top++
	s.s[s.top] = x
}

func (s *Stack) Pop() int {
	s.top--
	return s.s[s.top+1]
}

func main() {
	var l string
	fmt.Scan(&l)

	s := new(Stack)
	s.s = make([]int, len(l))
	s.top = -1

	sum := 0
	for i, v := range l {
		fmt.Println(s.s, s.top)
		if v == '\\' {
			s.Push(i)
		} else if v == '/' && s.top >= 0 {
			j := s.Pop()
			sum += i - j
		}
	}

	fmt.Println(sum)
}
