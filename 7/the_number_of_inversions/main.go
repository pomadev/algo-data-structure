package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MAX      = 200000
	SENTINEL = 2000000000
)

var l, r = make([]int, MAX/2+2), make([]int, MAX/2+2)

func merge(a []int, left, mid, right int) int {
	n1, n2 := mid-left, right-mid
	for i := 0; i < n1; i++ {
		l[i] = a[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = a[mid+i]
	}
	l[n1], r[n2] = SENTINEL, SENTINEL
	i, j, cnt := 0, 0, 0
	for k := left; k < right; k++ {
		if l[i] <= r[j] {
			a[k] = l[i]
			i++
		} else {
			a[k] = r[j]
			j++
			cnt += n1 - i
		}
	}
	return cnt
}

func mergeSort(a []int, left, right int) int {
	if left+1 < right {
		mid := (left + right) / 2
		v1 := mergeSort(a, left, mid)
		v2 := mergeSort(a, mid, right)
		v3 := merge(a, left, mid, right)
		return v1 + v2 + v3
	} else {
		return 0
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}

	ans := mergeSort(a, 0, n)
	fmt.Println(ans)
}
