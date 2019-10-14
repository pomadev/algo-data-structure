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

type Pair struct {
	first, second int
}

type Stack2 struct {
	p   []Pair
	top int
}

func (s2 *Stack2) Push(p Pair) {
	s2.top++
	s2.p[s2.top] = p
}

func (s2 *Stack2) Pop() Pair {
	s2.top--
	return s2.p[s2.top+1]
}

func main() {
	var l string
	fmt.Scan(&l)

	s := new(Stack)
	s.s = make([]int, len(l))
	s.top = -1
	s2 := new(Stack2)
	s2.p = make([]Pair, len(l))
	s2.top = -1

	sum := 0
	for i, v := range l {
		if v == '\\' {
			s.Push(i)
		} else if v == '/' && s.top >= 0 {
			j := s.Pop()
			sum += i - j
			a := i - j
			for s2.top >= 0 && s2.p[s2.top].first > j {
				a += s2.Pop().second
			}
			s2.Push(Pair{j, a})
		}
	}

	fmt.Println(sum)
	var ans []int
	for s2.top >= 0 {
		ans = append(ans, s2.Pop().second)
	}
	fmt.Print(len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Print(" ", ans[i])
	}
	fmt.Println()
}
