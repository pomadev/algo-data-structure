package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	h int
	a []int
)

func maxHeapify(i int) {
	l, r := 2*i, 2*i+1
	var largest int
	if l <= h && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= h && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(largest)
	}
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	h = nextInt()

	a = make([]int, h+1)
	for i := 1; i <= h; i++ {
		a[i] = nextInt()
	}

	for i := h / 2; i >= 1; i-- {
		maxHeapify(i)
	}

	for i := 1; i <= h; i++ {
		fmt.Print(" ", a[i])
	}
	fmt.Println()
}
