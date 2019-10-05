package main

import "fmt"

type Card struct {
	suit  string
	value int
}

func bubble(a []Card) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j >= i+1; j-- {
			if a[j].value < a[j-1].value {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func selection(a []Card) {
	for i := 0; i < len(a); i++ {
		minj := i
		for j := i; j < len(a); j++ {
			if a[j].value < a[minj].value {
				minj = j
			}
		}
		a[i], a[minj] = a[minj], a[i]
	}
}

func print(a []Card) {
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v.suit, v.value)
	}
	fmt.Println()
}

func isStable(c1, c2 []Card) bool {
	for i := 0; i < len(c1); i++ {
		if c1[i].suit != c2[i].suit {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]Card, n)
	for i := range a {
		fmt.Scanf("%1s%1d", &a[i].suit, &a[i].value)
	}

	b := make([]Card, n)
	copy(b, a)

	bubble(a)
	selection(b)

	print(a)
	fmt.Println("Stable")
	print(b)
	if isStable(a, b) {
		fmt.Println("Stable")
		return
	}
	fmt.Println("Not stable")
}
