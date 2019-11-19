package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	h int
	a = make([]int, 21832059)
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
		a[largest], a[i] = a[i], a[largest]
		maxHeapify(largest)
	}
}

func insert(key int) {
	h++
	a[h] = -1
	heapIncreaseKey(h, key)
}

func heapIncreaseKey(i, key int) {
	if key < a[i] {
		return
	}
	a[i] = key
	for i > 1 && a[i/2] < a[i] {
		a[i], a[i/2] = a[i/2], a[i]
		i = i / 2
	}
}

func heapExtractMax() int {
	if h < 1 {
		return -1
	}
	max := a[1]
	a[1] = a[h]
	h--
	maxHeapify(1)
	return max
}

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	n, _ := strconv.Atoi(nextStr())
	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	for {
		switch nextStr() {
		case "insert":
			insert(nextInt())
		case "extract":
			fmt.Println(heapExtractMax())
		case "end":
			return
		}
	}
}
