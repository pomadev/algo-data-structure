package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const (
	INF = 2000000000
	MAX = 500000
)

var (
	count int
	// l, rをmerge関数の中で作ると時間オーバーする
	l, r = make([]int, MAX+2), make([]int, MAX+2)
)

func merge(a []int, left, mid, right int) {
	n1, n2 := mid-left, right-mid
	for i := 0; i < n1; i++ {
		l[i] = a[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = a[mid+i]
	}
	l[n1], r[n2] = INF, INF
	i, j := 0, 0
	for k := left; k < right; k++ {
		count++
		if l[i] <= r[j] {
			a[k] = l[i]
			i++
		} else {
			a[k] = r[j]
			j++
		}
	}
}

func mergeSort(a []int, n, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(a, n, left, mid)
		mergeSort(a, n, mid, right)
		merge(a, left, mid, right)
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
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}

	mergeSort(s, n, 0, n)

	var b bytes.Buffer
	for i := 0; i < n; i++ {
		if i > 0 {
			b.Write([]byte(" "))
		}
		b.Write([]byte(fmt.Sprint(s[i])))
	}
	b.Write([]byte("\n"))
	b.Write([]byte(fmt.Sprintln(count)))
	b.WriteTo(os.Stdout)
}
