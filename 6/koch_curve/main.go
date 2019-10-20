package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func koch(d int, a, b Point) {
	if d == 0 {
		return
	}

	var s, t, u Point

	s.x = (2.0*a.x + 1.0*b.x) / 3.0
	s.y = (2.0*a.y + 1.0*b.y) / 3.0
	t.x = (1.0*a.x + 2.0*b.x) / 3.0
	t.y = (1.0*a.y + 2.0*b.y) / 3.0
	u.x = (t.x-s.x)*math.Cos(math.Pi/3) - (t.y-s.y)*math.Sin(math.Pi/3) + s.x
	u.y = (t.x-s.x)*math.Sin(math.Pi/3) + (t.y-s.y)*math.Cos(math.Pi/3) + s.y

	koch(d-1, a, s)
	fmt.Printf("%.8f %.8f\n", s.x, s.y)
	koch(d-1, s, u)
	fmt.Printf("%.8f %.8f\n", u.x, u.y)
	koch(d-1, u, t)
	fmt.Printf("%.8f %.8f\n", t.x, t.y)
	koch(d-1, t, b)
}

func main() {
	var n int
	fmt.Scan(&n)

	var a, b Point
	a.x = 0.0
	a.y = 0.0
	b.x = 100.0
	b.y = 0.0

	fmt.Printf("%.8f %.8f\n", a.x, a.y)
	koch(n, a, b)
	fmt.Printf("%.8f %.8f\n", b.x, b.y)
}
