package main

import "fmt"

type process struct {
	name string
	time int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var (
	head, tail, n int
	p             []process
)

func enqueue(x process) {
	p[tail] = x
	// n + 1 -> 配列のサイズ
	tail = (tail + 1) % (n + 1)
}

func dequeue() process {
	x := p[head]
	// n + 1 -> 配列のサイズ
	head = (head + 1) % (n + 1)
	return x
}

func main() {
	var t int
	fmt.Scan(&n, &t)

	p = make([]process, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].name, &p[i].time)
	}
	head, tail = 0, n

	elaps := 0
	for head != tail {
		u := dequeue()
		c := min(t, u.time)
		u.time -= c
		elaps += c
		if u.time > 0 {
			enqueue(u)
			continue
		}
		fmt.Printf("%s %d\n", u.name, elaps)
	}
}
