package main

import "fmt"

const (
	MAX      = 100000
	SENTINEL = 2000000000
)

type Card struct {
	suit  string
	value int
}

var l, r = make([]Card, MAX/2+2), make([]Card, MAX/2+2)

func merge(a []Card, left, mid, right int) {
	n1, n2 := mid-left, right-mid
	for i := 0; i < n1; i++ {
		l[i] = a[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = a[mid+i]
	}
	l[n1].value, r[n2].value = SENTINEL, SENTINEL
	i, j := 0, 0
	for k := left; k < right; k++ {
		if l[i].value <= r[j].value {
			a[k] = l[i]
			i++
		} else {
			a[k] = r[j]
			j++
		}
	}
}

func mergeSort(a []Card, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(a, left, mid)
		mergeSort(a, mid, right)
		merge(a, left, mid, right)
	}
}

func partition(a []Card, p, r int) int {
	x := a[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if a[j].value <= x.value {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func quickSort(a []Card, p, r int) {
	if p < r {
		q := partition(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q+1, r)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]Card, n)
	for i := range a {
		fmt.Scan(&a[i].suit, &a[i].value)
	}
	b := make([]Card, n)
	copy(b, a)

	quickSort(a, 0, n-1)
	mergeSort(b, 0, n)

	isStable := true
	for i := 0; i < n; i++ {
		if a[i].suit != b[i].suit {
			isStable = false
			break
		}
	}
	if isStable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
	for _, v := range a {
		fmt.Println(v.suit, v.value)
	}
}
